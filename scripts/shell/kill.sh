if [ -r "../services.pid" ]; then
  while IFS= read -r line; do
    pid=$(echo "$line" | awk -F: '{print $2}')
    if [[ "$pid" =~ ^[0-9]+$ ]]; then
      kill "$pid"
      if [ $? -eq 0 ]; then
        echo "Process with PID $pid has been killed."
      else
        echo "failed to kill process with PID $pid."
      fi
    fi
  done < "../services.pid"
else
  echo "File not found or not readable."
fi