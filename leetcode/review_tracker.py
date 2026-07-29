import re
import json
import os
from datetime import datetime, timedelta

LOG = "rapid-fire-log.md"
SCHEDULE = "review_schedule.json"

INTERVALS = [1, 3, 7, 14, 30]  # days after initial log

def parse_entries():
    with open(LOG) as f:
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

    result = []
    for e in entries:
        result.append({
            "id": int(e[0]),
            "title": e[1].strip(),
            "pattern": e[5].strip(),
            "insight": e[6].strip()
        })
    return result

def load_schedule():
    if os.path.exists(SCHEDULE):
        with open(SCHEDULE) as f:
            return json.load(f)
    return {}

def save_schedule(sched):
    with open(SCHEDULE, "w") as f:
        json.dump(sched, f, indent=2)

def init_schedule(entries, sched):
    today = datetime.now().replace(hour=0, minute=0, second=0, microsecond=0)

    for e in entries:
        pid = str(e["id"])
        if pid in sched:
            continue

        gap_analysis = "Pattern Coverage Gap Analysis"
        with open(LOG) as f:
            content = f.read()
            gap_idx = content.find(gap_analysis)

        log_date = today
        sched[pid] = {
            "id": e["id"],
            "title": e["title"],
            "pattern": e["pattern"],
            "insight": e["insight"],
            "created": log_date.strftime("%Y-%m-%d"),
            "reviews": [(log_date + timedelta(days=d)).strftime("%Y-%m-%d") for d in INTERVALS],
            "done": []
        }
    save_schedule(sched)
    return sched

def show_due(sched):
    today = datetime.now().strftime("%Y-%m-%d")
    due = []
    for pid, entry in sched.items():
        if today in entry["reviews"] and today not in entry["done"]:
            due.append(entry)

    print(f"=== Review Queue for {today} ===\n")
    if not due:
        print("Nothing due today. Rest or review flagged problems.\n")
    else:
        for e in sorted(due, key=lambda x: x["id"]):
            print(f"LC{e['id']} — {e['title']}")
            print(f"  Pattern: {e['pattern']}")
            print(f"  Key insight: {e['insight']}")
            print()

    print(f"Total due: {len(due)}")
    return due

def mark_done(pid):
    sched = load_schedule()
    today = datetime.now().strftime("%Y-%m-%d")
    pid = str(pid)
    if pid in sched and today in sched[pid]["reviews"]:
        sched[pid]["done"].append(today)
        save_schedule(sched)
        print(f"LC{pid} marked reviewed for {today}")

def stats(sched):
    today = datetime.now().strftime("%Y-%m-%d")
    total = len(sched)
    due_today = sum(1 for e in sched.values() if today in e["reviews"] and today not in e["done"])
    completed = sum(1 for e in sched.values() if today in e["done"])

    print(f"Total problems: {total}")
    print(f"Due today: {due_today}")
    print(f"Completed today: {completed}")

    all_reviews = set()
    done_reviews = set()
    for e in sched.values():
        for r in e["reviews"]:
            all_reviews.add((e["id"], r))
        for d in e["done"]:
            done_reviews.add((e["id"], d))
    print(f"Total reviews scheduled: {len(all_reviews)}")
    print(f"Total reviews completed: {len(done_reviews)}")
    print(f"Completion rate: {len(done_reviews)/len(all_reviews)*100:.0f}%" if all_reviews else "N/A")

if __name__ == "__main__":
    import sys
    os.chdir(os.path.dirname(os.path.abspath(__file__)))

    entries = parse_entries()
    sched = load_schedule()

    if not sched:
        print(f"Initializing schedule for {len(entries)} problems...")
        sched = init_schedule(entries, sched)
        print("Done.\n")

    if len(sys.argv) > 1:
        if sys.argv[1] == "done":
            if len(sys.argv) > 2:
                mark_done(sys.argv[2])
            else:
                print("Usage: python review_tracker.py done <LC_ID>")
        elif sys.argv[1] == "stats":
            stats(sched)
        elif sys.argv[1] == "reinit":
            sched = init_schedule(entries, sched)
            print(f"Re-initialized. {len(sched)} problems.")
        else:
            print("Commands: (no args) = show due, done <id>, stats, reinit")
    else:
        show_due(sched)
