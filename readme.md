# **File Modification Tracker in Go**

## **Overview**

This project implements a **File Modification Tracker** service for Windows using **Go**. The service runs as a Windows daemon, monitors file modifications in a specified directory, and integrates with **osquery** for system monitoring. It provides both a simple native Windows UI for service control and an HTTP interface for health checks and logs retrieval. Configuration is managed using **viper** and validated with **validator**.

---

## **Features**
- **Windows Service (Daemon)**: Runs in the background, tracking file modifications.
- **osquery Integration**: Fetches file modification stats and other system metrics.
- **Multi-threaded Design**: Utilizes a worker thread and a timer thread.
- **Native Windows UI**: Start/stop service and view logs.
- **HTTP API**: Exposes endpoints for health checks and log retrieval.
- **Remote API Integration**: Sends file modification stats to a remote server.
- **Configuration Management**: Managed through a `config.yaml` file using viper.
- **Packaging**: Packaged into a `.msi` installer for Windows.

---

## **Project Structure**

```
.
├── cmd/
│   └── main.go             # Main entry point for the service
├── internal/
│   ├── api/
│   │   └── client.go       # API client for remote reporting
│   ├── config/
│   │   └── config.go       # Configuration loading with viper
│   ├── service/
│   │   ├── worker_thread.go  # Worker thread implementation
│   │   └── timer_thread.go   # Timer thread for file modification tracking
│   ├── osquery/
│   │   └── osquery.go      # osquery integration for file tracking
│   ├── ui/
│   │   └── ui.go           # Native Windows UI for controlling the service
├── pkg/
│   └── logs/
│       └── logger.go       # Logger setup
├── resources/
│   └── config.yaml         # Configuration file
├── tests/
│   ├── service_test.go     # Unit tests for service functionality
│   └── http_test.go        # Unit tests for HTTP endpoints
├── installer/
│   └── installer.wxs       # WiX toolset installer script for .msi packaging
└── README.md               # Project documentation
```

---

## **Requirements**

- **Go 1.16+**: Ensure you have the latest Go version installed.
- **osquery**: Download and install osquery [here](https://osquery.io/downloads).
- **Windows 10/11**: The service is designed specifically for Windows.
- **WiX Toolset**: To package the project into an MSI installer. Download [here](https://wixtoolset.org/).

---

## **Installation**

### **1. Clone the repository**:
```bash
git clone https://github.com/marveloper-8/file-modification-tracker-go.git
cd file-modification-tracker-go
```

### **2. Install dependencies**:
```bash
go mod tidy
```

### **3. Build the service**:
If you are building directly on a Windows machine:
```bash
go build -o file_modification_tracker.exe ./cmd/main.go
```

If you are cross-compiling from a non-Windows platform:
```bash
GOOS=windows GOARCH=amd64 go build -o file_modification_tracker.exe ./cmd/main.go
```

### **4. Package into an MSI installer**:
First, install the **WiX Toolset** and ensure that `candle.exe` and `light.exe` are in your system’s PATH.

Then, run the following command to package the service into an `.msi` installer:

```bash
candle installer/installer.wxs
light installer.wixobj -o file_modification_tracker.msi
```

### **5. Install the service**:
Run the `.msi` file you just created, which will install and register the service as a Windows service.

---

## **Configuration**

### **Configuration File (config.yaml)**

The service is configured using a `config.yaml` file, located in the `resources/` directory. The file looks like this:

```yaml
directory_to_monitor: "C:/path/to/your/directory"
check_frequency: 1 # Frequency of checks in minutes
api_endpoint: "https://your-api-endpoint.com"
```

- **directory_to_monitor**: The directory where file changes will be tracked.
- **check_frequency**: The frequency (in minutes) for checking file modifications.
- **api_endpoint**: The remote API to which file stats will be sent.

---

## **Usage**

### **Starting the Service**

The service runs automatically once installed, but you can also manage it via the Windows `sc` utility.

To start the service manually:

```bash
sc start FileModificationTracker
```

To stop the service:

```bash
sc stop FileModificationTracker
```

### **Using the Native UI**

You can interact with the service using the native Windows UI, which allows you to:
- Start/Stop the service
- View logs

Simply run the `file_modification_tracker.exe` to open the UI.

---

## **HTTP API Endpoints**

### **1. Health Check Endpoint**:

- **URL**: `http://localhost:8080/health`
- **Method**: `GET`
- **Description**: Returns the health status of the worker and timer threads.

### **2. Logs Retrieval Endpoint**:

- **URL**: `http://localhost:8080/logs`
- **Method**: `GET`
- **Description**: Retrieves logs of file modification activities.

---

## **Remote API Integration**

The service sends file modification stats every minute to the configured remote API endpoint. This is handled via the `http.Client` in the `internal/api/client.go` file.

---

## **Unit Tests**

Unit tests are provided for core functionalities, including service threads, API client, and HTTP endpoints.

To run the tests:

```bash
go test ./...
```

---

## **Uninstallation**

To uninstall the service:

1. Open **Control Panel** > **Programs** > **Uninstall a Program**.
2. Locate **File Modification Tracker** and uninstall it.

Alternatively, you can uninstall via the command line:

```bash
msiexec /x file_modification_tracker.msi
```

---

## **Contributing**

Feel free to fork the repository, submit issues, and contribute to the project by opening pull requests.

---

## **License**

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## **Authors**

- **Joshua Samuel** - Fullstack Cloud Software Engineer (GitHub: [marveloper-8](https://github.com/marveloper-8))

---

### **Happy tracking!**