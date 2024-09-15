# File Modification Tracker

### Project Description
The **File Modification Tracker** is a Windows service built with Go that monitors a specified directory for file modifications. It tracks changes using `osquery` for system monitoring and logs the results. The service runs as a background daemon and provides several functionalities:

- Tracks file modifications in a given directory using `osquery`.
- Manages configuration using `viper` and validates it with `validator`.
- Offers HTTP endpoints to interact with the service for health checks, logs retrieval, and worker thread management.
- Provides a native Windows UI built with `Fyne` to allow users to start/stop the service and view logs.
- Packages the service as an `.msi` installer for Windows.

---

### Features
- **Daemon**: Runs in the background as a Windows service with worker and timer threads.
- **File Monitoring**: Uses `osquery` to retrieve file modification statistics.
- **HTTP API**: Provides endpoints for submitting commands to the worker thread, checking the health of the service, and retrieving logs.
- **UI Component**: A native Windows UI to interact with the service.
- **Logging**: Logs file modification statistics and service activity.
- **Configurable**: Manages settings through `viper` and validates with `validator`.
- **MSI Packaging**: The service is packaged as an `.msi` file for easy installation and uninstallation on Windows.

---

### Prerequisites
Ensure you have the following installed on your system:
- **Go (1.18 or later)**: [Download Go](https://golang.org/dl/)
- **WiX Toolset** (for building MSI packages): [WiX Toolset](https://wixtoolset.org/)
- **osquery**: [Download osquery](https://osquery.io/downloads/official/windows)

---

### Project Structure
```bash
file-modification-tracker/
├── cmd/
│   └── service/
│       └── main.go            # Main service code
├── config/
│   └── config.go              # Configuration handling
├── daemon/
│   ├── daemon.go              # Daemon management
│   ├── worker.go              # Worker thread
├── http/
│   └── server.go              # HTTP server
├── logs/
│   └── logger.go              # Logging functionality
├── osquery/
│   └── osquery.go             # Osquery integration
├── ui/
│   └── ui.go                  # Windows UI using Fyne
├── tests/
│   ├── config_test.go         # Unit tests for configuration
│   ├── worker_test.go         # Unit tests for worker thread
│   └── timer_test.go          # Unit tests for timer thread
├── config.yaml                # Configuration file (YAML)
├── service.wxs                # WiX configuration for MSI packaging
├── go.mod                     # Go module file
├── go.sum                     # Go dependencies
└── README.md                  # Project documentation
```

---

### Configuration

The configuration file `config.yaml` is used to manage the service settings. The file should be placed in the root directory and contain the following settings:

```yaml
directory: "C:/path/to/monitor"  # Directory to monitor
check_freq: 60                   # Frequency (in seconds) of file modification checks
remote_api: "https://api.example.com/collect"  # Remote API to send stats
```

You can customize the directory, frequency of checks, and API endpoint for reporting.

---

### Build Instructions

#### 1. **Clone the Repository**

```bash
git clone https://github.com/your-username/file-modification-tracker.git
cd file-modification-tracker
```

#### 2. **Install Dependencies**

```bash
go mod tidy
```

#### 3. **Install Required Go Libraries**

```bash
# Install viper for configuration management
go get github.com/spf13/viper

# Install validator for configuration validation
go get github.com/go-playground/validator/v10

# Install Fyne for UI
go get fyne.io/fyne/v2

# Install testify for unit testing
go get github.com/stretchr/testify
```

#### 4. **Build the Go Binary**

To build the service for Windows:

```bash
GOOS=windows GOARCH=amd64 go build -o file-modification-tracker.exe ./cmd/service
```

#### 5. **Create MSI Installer**

Make sure you have WiX Toolset installed, then use the following commands to create an `.msi` installer:

```bash
# Compile the WiX source file
candle service.wxs

# Link and create the MSI package
light -out service.msi service.wixobj
```

#### 6. **Install the MSI Package**

Once the `.msi` is created, you can install the service using the following command:

```bash
msiexec /i service.msi
```

#### 7. **Run Unit Tests**

To run the unit tests for the core functionalities (configuration, worker, timer):

```bash
go test ./tests
```

---

### Usage Instructions

1. **Start the Service**:
   - Run the installed service through Windows Service Manager, or execute the binary directly:
   
   ```bash
   ./file-modification-tracker.exe
   ```

2. **Interact with the Service**:
   - Use the native Windows UI (built using Fyne) to start/stop the service and view logs.
   - Or, access the HTTP endpoints:
     - **Submit Commands to Worker Thread**:
       ```bash
       curl -X POST http://localhost:8080/commands -d '{"commands": ["echo Hello"]}'
       ```
     - **Check Service Health**:
       ```bash
       curl http://localhost:8080/health
       ```
     - **Retrieve Logs**:
       ```bash
       curl http://localhost:8080/logs
       ```

---

### Uninstallation

To uninstall the service, run the following command:

```bash
msiexec /x service.msi
```

---

### License
This project is licensed under the MIT License.