#!/usr/bin/env python3
"""Exercise smoke EXIT cleanup with fake tools; never starts containers."""
import os
from pathlib import Path
import subprocess
import tempfile

root = Path(__file__).resolve().parents[1]
with tempfile.TemporaryDirectory(prefix="i3d-cleanup-") as temp:
    directory = Path(temp)
    docker = directory / "docker"
    docker.write_text("""#!/bin/sh
printf '%s\\n' "$*" >> "$MOCK_CALL_LOG"
case " $* " in
  *" up "*) exit "$MOCK_UP_EXIT" ;;
  *" logs "*) exit "$MOCK_LOGS_EXIT" ;;
  *" down "*) exit "$MOCK_DOWN_EXIT" ;;
  *" port "*) printf '127.0.0.1:12345\\n' ;;
  *) exit 99 ;;
esac
""")
    docker.chmod(0o700)
    for name, content in {"python3": "#!/bin/sh\nexit 0\n",
                          "go": '#!/bin/sh\nexit "$MOCK_CLIENT_EXIT"\n'}.items():
        tool = directory / name
        tool.write_text(content)
        tool.chmod(0o700)
    for up, logs, down, client in [(7, 0, 0, 0), (7, 9, 0, 0), (7, 0, 11, 0),
                                   (0, 9, 11, 5), (0, 0, 11, 0), (0, 0, 0, 0)]:
        log = directory / "calls"
        log.write_text("")
        env = dict(os.environ, PATH=str(directory) + os.pathsep + os.environ["PATH"],
                   MOCK_CALL_LOG=str(log), MOCK_UP_EXIT=str(up), MOCK_LOGS_EXIT=str(logs),
                   MOCK_DOWN_EXIT=str(down), MOCK_CLIENT_EXIT=str(client))
        result = subprocess.run(["sh", str(root / "scripts/smoke.sh")], env=env,
                                capture_output=True, text=True)
        expected = up or client
        assert result.returncode == expected, (up, logs, down, client, result.returncode, result.stderr)
        calls = log.read_text().splitlines()
        assert sum(" down " in " " + line + " " for line in calls) == 1, calls
        assert any(" logs " in " " + line + " " for line in calls) == bool(expected), calls
print("Smoke cleanup preserves the scenario exit status and always attempts teardown.")
