# -*- coding: utf-8 -*-
"""Fill 12-go-network/legacy-topic-index/*.md with 内容大纲. Run from repo root: python 12-go-network/tools/apply_outlines.py"""
from __future__ import annotations

import json
from pathlib import Path

# 大纲骨架仍写在「按技术主题」的 legacy 目录下（与 outlines/*.json 中 rel 路径一致）
ROOT = Path(__file__).resolve().parents[1] / "legacy-topic-index"


def write_md(rel: str, title: str, folder_zh: str, bullets: list[str]) -> None:
    p = ROOT / rel
    lines = [
        f"# {title}",
        "",
        f"> **{folder_zh}**",
        "",
        "## 内容大纲",
        "",
    ]
    for b in bullets:
        lines.append(f"- {b}")
    lines.extend(["", "## 正文", "", "（待补充）", ""])
    p.write_text("\n".join(lines), encoding="utf-8")


def main() -> None:
    out_dir = Path(__file__).with_name("outlines")
    payload: list = []
    for f in sorted(out_dir.glob("*.json")):
        payload.extend(json.loads(f.read_text(encoding="utf-8")))
    for item in payload:
        write_md(item["rel"], item["title"], item["folder"], item["bullets"])
    print(f"wrote {len(payload)} files under {ROOT}")


if __name__ == "__main__":
    main()
