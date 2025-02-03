#!/bin/bash

if [ $# -ne 4 ]; then
    echo "Usage: $0 <OS> <ARCH> <install|update> <API_KEY>"
    echo "Example: $0 Linux amd64 install your_api_key"
    exit 1
fi

OS=$1
ARCH=$2
ACTION=$3
API_KEY=$4
FILE_NAME="soliduptime-agent_${OS}_${ARCH}.tar.gz"

DOWNLOAD_URL="https://github.com/WebduoNederland/soliduptime-agent/releases/latest/download/$FILE_NAME"

echo "Downloading $FILE_NAME from $DOWNLOAD_URL ..."
tar -xzf <(curl -Ls "$DOWNLOAD_URL") -C /usr/local/bin
echo "Finished!"

echo "Change agent file permissions..."
chmod +x /usr/local/bin/soliduptime-agent
echo "Finished!"

if [[ "$OS" == "linux" ]]; then
    SERVICE_FILE="/etc/systemd/system/soliduptime-agent.service"
    
    if [[ "$ACTION" == "install" ]]; then
        echo "Creating systemd service..."
        echo "[Unit]
        Description=SolidUptime Agent
        After=network.target

        [Service]
        ExecStart=/usr/local/bin/soliduptime-agent --api-key $API_KEY
        Restart=always
        User=root

        [Install]
        WantedBy=multi-user.target" | sudo tee "$SERVICE_FILE" > /dev/null
        echo "Finished!"

        echo "Enabling and starting the service..."
        sudo systemctl daemon-reload
        sudo systemctl enable soliduptime-agent
        sudo systemctl start soliduptime-agent
        echo "Finished!"

    elif [[ "$ACTION" == "update" ]]; then
        echo "Restarting service to apply update..."
        sudo systemctl restart soliduptime-agent
        echo "Finished!"
    else
        echo "Invalid action: $ACTION. Use 'install' or 'update'."
        exit 1
    fi
fi

echo "SolidUptime agent $ACTION completed!"
