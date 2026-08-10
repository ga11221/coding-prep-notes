import re
import json
import os
from datetime import datetime, timedelta

REVIEW_DIR = os.path.dirname(os.path.abspath(__file__))
SCHEDULE_FILE = os.path.join(REVIEW_DIR, "review_schedule.json")

INTERVALS = [1, 3, 7, 14, 30]

def load():
    if os.path.exists(SCHEDULE_FILE):
        with open(SCHEDULE_FILE) as f:
            return json.load(f)
    return {}

def save(sched):
    with open(SCHEDULE_FILE, "w") as f:
        json.dump(sched, f, indent=2)

def add(topic, item_id, title, detail, created=None):
    sched = load()
    key = f"{topic}:{item_id}"
    if key in sched:
        print(f"Exists: {key}")
        return

    if created is None:
        created = datetime.now()
    elif isinstance(created, str):
        created = datetime.strptime(created, "%Y-%m-%d")
    else:
        created = created.replace(hour=0, minute=0, second=0, microsecond=0)

    sched[key] = {
        "topic": topic,
        "id": item_id,
        "title": title,
        "detail": detail,
        "created": created.strftime("%Y-%m-%d"),
        "reviews": [(created + timedelta(days=d)).strftime("%Y-%m-%d") for d in INTERVALS],
        "done": []
    }
    save(sched)
    print(f"Added: {key}")

def _last_done(entry, today):
    done = [d for d in entry.get("done", []) if d <= today]
    return max(done) if done else None

def _missed(entry, today):
    """Scheduled reviews (<= today) not covered by the most recent review.

    One review clears all prior missed slots: only dates after the last done
    date count as missed.
    """
    last = _last_done(entry, today)
    return sum(1 for d in entry["reviews"] if d <= today and (last is None or d > last))

def due(sched=None):
    if sched is None:
        sched = load()
    today = datetime.now().strftime("%Y-%m-%d")
    items = []
    for key, entry in sched.items():
        if today in entry["reviews"] and today not in entry["done"]:
            items.append((key, entry))
    return sorted(items, key=lambda x: x[1]["topic"] + str(x[1]["id"]))

def overdue(sched=None, limit=None):
    """Items with scheduled review slots (<= today) not covered by done.

    A review covers all missed slots up to its date (see mark_done).
    """
    if sched is None:
        sched = load()
    today = datetime.now().strftime("%Y-%m-%d")
    items = []
    for key, entry in sched.items():
        missed = _missed(entry, today)
        if missed > 0:
            last = [d for d in entry["reviews"] if d <= today][-1]
            items.append((missed, key, entry, last))
    items.sort(key=lambda x: (-x[0], x[3]))
    if limit:
        items = items[:limit]
    return items

def show_overdue(limit=None):
    all_items = overdue(limit=None)
    total_missed = sum(i[0] for i in all_items)
    items = all_items if limit is None else all_items[:limit]
    today = datetime.now().strftime("%Y-%m-%d")
    print(f"=== Overdue backlog (total {len(all_items)} items, {total_missed} missed slots) ===")
    print(f"Showing {len(items)}. Re-run with a count to see more (e.g. 'overdue 30').\n")

    if not items:
        print("Nothing overdue. You're caught up!\n")
        return

    by_topic = {}
    for missed, key, entry, last in items:
        by_topic.setdefault(entry["topic"], []).append((missed, key, entry, last))

    for topic in sorted(by_topic):
        entries = by_topic[topic]
        print(f"[{topic}] ({len(entries)})")
        for missed, key, entry, last in entries:
            print(f"  {entry['id']} — {entry['title']}   (missed {missed}, last scheduled {last})")
        print()

    print("Process: review an item, then: done <topic> <id>. One review clears its missed slots.")

def show_due():
    items = due()
    today = datetime.now().strftime("%Y-%m-%d")
    print(f"=== Review Queue for {today} ===\n")

    if not items:
        print("Nothing due today.\n")
        return

    by_topic = {}
    for key, entry in items:
        by_topic.setdefault(entry["topic"], []).append((key, entry))

    total = 0
    for topic in sorted(by_topic):
        entries = by_topic[topic]
        print(f"[{topic}] ({len(entries)})")
        for key, entry in entries:
            print(f"  {entry['id']} — {entry['title']}")
            print(f"    {entry['detail']}")
        print()
        total += len(entries)

    print(f"Total due: {total}")

def mark_done(topic, item_id):
    sched = load()
    key = f"{topic}:{item_id}"
    today = datetime.now().strftime("%Y-%m-%d")
    if key not in sched:
        print(f"Not found: {key}")
        return
    entry = sched[key]
    if _missed(entry, today) == 0 and today not in entry["reviews"]:
        print(f"Not due or overdue: {key}")
        return
    if today not in entry["done"]:
        entry["done"].append(today)
    save(sched)
    print(f"Marked done: {key}")

def stats():
    sched = load()
    today = datetime.now().strftime("%Y-%m-%d")
    total = len(sched)

    by_topic = {}
    for key, entry in sched.items():
        by_topic.setdefault(entry["topic"], 0)
        by_topic[entry["topic"]] += 1

    print("=== Review Stats ===\n")
    for topic, count in sorted(by_topic.items()):
        due_today = sum(1 for k, e in sched.items() if e["topic"] == topic and today in e["reviews"] and today not in e["done"])
        ov = sum(1 for k, e in sched.items() if e["topic"] == topic and _missed(e, today) > 0)
        print(f"{topic}: {count} items, {due_today} due today, {ov} overdue")

    all_rev = sum(len(e["reviews"]) for e in sched.values())
    done_rev = sum(len(e["done"]) for e in sched.values())
    ov_total = sum(1 for e in sched.values() if _missed(e, today) > 0)
    missed_total = sum(_missed(e, today) for e in sched.values())
    print(f"\nTotal: {total} items, {all_rev} reviews scheduled, {done_rev} completed")
    print(f"Overdue: {ov_total} items, {missed_total} missed slots")
    if all_rev:
        print(f"Completion rate: {done_rev/all_rev*100:.0f}%")

def list_topics():
    sched = load()
    by_topic = {}
    for key, entry in sched.items():
        by_topic.setdefault(entry["topic"], []).append((key, entry))

    print("=== Topics ===\n")
    for topic in sorted(by_topic):
        entries = by_topic[topic]
        print(f"  {topic}: {len(entries)} items")
    print(f"\n  Total: {len(sched)} items")

def list_topic(topic):
    sched = load()
    items = [(k, v) for k, v in sched.items() if v["topic"].lower() == topic.lower()]
    if not items:
        print(f"Topic '{topic}' not found")
        return

    print(f"=== {topic.upper()} ({len(items)}) ===\n")
    for key, entry in sorted(items, key=lambda x: str(x[1]["id"])):
        print(f"  {entry['id']} — {entry['title']}")
        print(f"    Created: {entry['created']}")
        print(f"    Next reviews: {', '.join(r for r in entry['reviews'] if r not in entry.get('done', []))}")
        print()

def reinit_lc():
    """(Re)import all LC problems from rapid-fire-log.md into schedule"""
    log_path = os.path.join(REVIEW_DIR, "leetcode", "rapid-fire-log.md")
    if not os.path.exists(log_path):
        print("rapid-fire-log.md not found")
        return

    with open(log_path) as f:
        content = f.read()

    entries = re.findall(
        r'## LC(\d+) — (.+?)\n'
        r'\*\*User approach:\*\* (.+?)\n'
        r'\*\*Refined approach:\*\* (.+?)\n'
        r'\*\*Time:\*\* (.+?)\n'
        r'\*\*Pattern:\*\* (.+?)\n'
        r'\*\*Key insight:\*\* (.+?)\n'
        r'\*\*Discrete Math:\*\* (.+?)(?:\n|$)',
        content, re.DOTALL
    )

    # Find gap analysis section to split old vs new
    gap_idx = content.find("Pattern Coverage Gap Analysis")
    if gap_idx < 0:
        gap_idx = len(content)

    before_gap = content[:gap_idx]
    new_ids = set()
    old_ids = set()
    for m in re.finditer(r'## LC(\d+)', before_gap):
        pos = m.start()
        if "Jul 28" in before_gap[max(0,pos-200):pos]:
            new_ids.add(m.group(1))
        else:
            old_ids.add(m.group(1))

    old_date = "2026-07-25"
    new_date = "2026-07-28"

    count = 0
    for e in entries:
        pid = e[0]
        title = e[1].strip()
        detail = f"Pattern: {e[5].strip()}"
        created = new_date if pid in new_ids else old_date
        add("LC", pid, title, detail, created=created)
        count += 1
    print(f"Imported {count} LC problems")

def repair():
    """Sort/dedupe review dates and drop any scheduled before creation.

    Fixes entries corrupted by manual editing (out-of-order review lists and
    review dates predating the item's creation date).
    """
    sched = load()
    changed = 0
    for key, entry in sched.items():
        created = entry["created"]
        rev = entry["reviews"]
        fixed = sorted(set(d for d in rev if d >= created))
        if fixed != rev:
            print(f"FIX {key}: {rev} -> {fixed}")
            entry["reviews"] = fixed
            changed += 1
        done = entry.get("done", [])
        dfixed = sorted(set(done))
        if dfixed != done:
            print(f"FIX done {key}: {done} -> {dfixed}")
            entry["done"] = dfixed
            changed += 1
    if changed:
        save(sched)
    print(f"Repaired {changed} entries")

if __name__ == "__main__":
    import sys

    if len(sys.argv) == 1:
        show_due()
    elif sys.argv[1] == "done" and len(sys.argv) >= 4:
        mark_done(sys.argv[2], sys.argv[3])
    elif sys.argv[1] == "stats":
        stats()
    elif sys.argv[1] == "overdue":
        n = int(sys.argv[2]) if len(sys.argv) >= 3 else None
        show_overdue(limit=n)
    elif sys.argv[1] == "repair":
        repair()
    elif sys.argv[1] == "reinit-lc":
        reinit_lc()
    elif sys.argv[1] == "add" and len(sys.argv) >= 5:
        add(sys.argv[2], sys.argv[3], sys.argv[4], sys.argv[5] if len(sys.argv) > 5 else "")
    elif sys.argv[1] == "list":
        if len(sys.argv) >= 3:
            list_topic(sys.argv[2])
        else:
            list_topics()
    else:
        print("Commands:")
        print("  (no args)          Show today's due items")
        print("  stats              Show stats")
        print("  list               List all topics")
        print("  list <topic>       List items in a topic")
        print("  overdue [N]        Show top N most-overdue items (default 10)")
        print("  done <topic> <id>  Mark item reviewed (works for overdue too)")
        print("  repair             Fix corrupted review dates (sort/dedupe/drop <created)")
        print("  reinit-lc          (Re)import all LC problems from log")
        print("  add <topic> <id> <title> [detail]  Add custom review item")
