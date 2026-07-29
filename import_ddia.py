import re
import json
import os
from datetime import datetime, timedelta

with open("ddia/notes.md") as f:
    text = f.read()

sched_file = "review_schedule.json"
sched = {}
if os.path.exists(sched_file):
    with open(sched_file) as f:
        sched = json.load(f)

INTERVALS = [1, 3, 7, 14, 30]

sections = [
    ("# DDIA Chapter 3", "2026-07-21"),
    ("# DDIA Chapter 7", "2026-07-21"),
    ("# DDIA Chapter 4", "2026-07-22"),
    ("# DDIA Chapter 11", "2026-07-22"),
    ("# Cumulative Quiz", "2026-07-22"),
    ("# Cumulative Quiz Round 2", "2026-07-24"),
    ("# DDIA Chapter 5", "2026-07-25"),
    ("### Multi-Leader", "2026-07-28"),
]

def find_date(pos):
    best_pos, best_date = -1, "2026-07-25"
    for header, date in sections:
        idx = text.rfind(header, 0, pos)
        if idx > best_pos:
            best_pos, best_date = idx, date
    return best_date

def make_reviews(created):
    d = datetime.strptime(created, "%Y-%m-%d")
    return sorted(set((d + timedelta(days=i)).strftime("%Y-%m-%d") for i in INTERVALS))

# Format 1: ### Q\d+: <question> \n\n **Answer:** <answer>
for m in re.finditer(
    r'### (Q\d*[^:]*: .+?)\n\n\*\*Answer:\*\* (.+?)(?:\n\n\*\*Feedback|\n\n---|\n\n\*\*Q|\n#)',
    text, re.DOTALL
):
    q = m.group(1).strip()
    a = m.group(2).strip()[:200]
    created = find_date(m.start())
    key = f"DDIA:{q[:50]}"
    if key not in sched:
        sched[key] = {
            "topic": "DDIA",
            "id": q[:50],
            "title": q,
            "detail": a,
            "created": created,
            "reviews": make_reviews(created),
            "done": []
        }

# Format 2: **Q: <question>** \n <answer>
for m in re.finditer(r'\*\*Q: (.+?)\*\*\n+(.+?)(?=\n\n\*\*Q:|\n\n##|\Z)', text, re.DOTALL):
    q = m.group(1).strip()
    a = m.group(2).strip()[:200]
    key = f"DDIA:{q[:50]}"
    if key not in sched:
        sched[key] = {
            "topic": "DDIA",
            "id": q[:50],
            "title": q,
            "detail": a,
            "created": "2026-07-28",
            "reviews": make_reviews("2026-07-28"),
            "done": []
        }

with open(sched_file, "w") as f:
    json.dump(sched, f, indent=2)

ddia = sum(1 for v in sched.values() if v["topic"] == "DDIA")
print(f"Done. DDIA: {ddia}, Total: {len(sched)}")
