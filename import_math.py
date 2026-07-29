import re, json, os
from datetime import datetime, timedelta

with open("math/rapid-fire-log.md") as f:
    text = f.read()

sched_file = "review_schedule.json"
with open(sched_file) as f:
    sched = json.load(f)

INTERVALS = [1, 3, 7, 14, 30]

def make_reviews(created):
    d = datetime.strptime(created, "%Y-%m-%d")
    return sorted(set((d + timedelta(days=i)).strftime("%Y-%m-%d") for i in INTERVALS))

# Simple approach: split on ## Q, parse each
blocks = re.split(r'\n(?=## Q\d+:)', text)
count = 0
for block in blocks:
    m = re.match(r'## Q(\d+): (.+?)\n', block)
    if not m:
        continue
    qnum = m.group(1)
    title = m.group(2).strip()

    qm = re.search(r'\*\*Q:\*\* (.+)', block)
    question = qm.group(1).strip() if qm else ""

    # Try Correction, then Refined answer, then Answer
    am = re.search(r'\*\*Correction:\*\* (.+)', block)
    if not am:
        am = re.search(r'\*\*Refined answer:\*\* (.+)', block)
    if not am:
        am = re.search(r'\*\*Answer:\*\* (.+)', block)
    answer = am.group(1).strip()[:200] if am else ""

    key = f"MATH:{title[:50]}"
    if key not in sched:
        sched[key] = {
            "topic": "MATH",
            "id": title[:50],
            "title": f"Q{qnum}: {title}",
            "detail": f"{question[:100]} → {answer[:100]}",
            "created": "2026-07-27",
            "reviews": make_reviews("2026-07-27"),
            "done": []
        }
        count += 1

with open(sched_file, "w") as f:
    json.dump(sched, f, indent=2)

print(f"Math items added: {count}, Total: {len(sched)}")
