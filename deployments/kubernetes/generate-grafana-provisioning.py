#!/usr/bin/env python3
"""
Generate grafana-datasources-configmap.yaml and grafana-dashboards-configmap.yaml
from observability/grafana/provisioning so the k8s Grafana deployment loads the
same datasources and dashboards as the docker-compose stack.

Usage: python3 generate-grafana-provisioning.py
Regenerates the two ConfigMap manifests in this directory. Re-run after editing
any file under observability/grafana/provisioning.
"""
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent.parent
PROV = ROOT / "observability" / "grafana" / "provisioning"
OUT = Path(__file__).resolve().parent


def yaml_scalar_block(text: str) -> str:
    lines = text.rstrip("\n").split("\n")
    return " |\n" + "\n".join("    " + ln for ln in lines) + "\n"


def build_configmap(name: str, files: dict) -> str:
    out = [
        "apiVersion: v1",
        "kind: ConfigMap",
        "metadata:",
        f"  name: {name}",
        "  namespace: payment-gateway",
        "data:",
    ]
    for key, content in files.items():
        out.append(f"  {key}:{yaml_scalar_block(content)}")
    return "\n".join(out) + "\n"


def main():
    datasource = (PROV / "datasources" / "datasource.yml").read_text()
    dashboards_cfg = (PROV / "dashboards" / "dashboards.yaml").read_text()

    dashboard_files = {}
    for p in sorted((PROV / "dashboards").glob("*.json")):
        dashboard_files[p.name] = p.read_text()

    files_map = {
        "datasources": {"datasource.yml": datasource},
        "dashboards": {**{"dashboards.yaml": dashboards_cfg}, **dashboard_files},
    }

    (OUT / "grafana-datasources-configmap.yaml").write_text(
        build_configmap("grafana-datasources", files_map["datasources"])
    )
    (OUT / "grafana-dashboards-configmap.yaml").write_text(
        build_configmap("grafana-dashboards", files_map["dashboards"])
    )

    print(f"Wrote grafana-datasources-configmap.yaml ({len(files_map['datasources'])} keys)")
    print(f"Wrote grafana-dashboards-configmap.yaml ({len(files_map['dashboards'])} keys)")
    for k in dashboard_files:
        print(f"  - {k}")


if __name__ == "__main__":
    main()
