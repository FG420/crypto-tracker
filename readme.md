# Crypto Tracker

## Overview
Crypto Tracker is a command-line application written in Go that enables users to manage and monitor the prices of various cryptocurrencies. The application offers features to view real-time market data of specific cryptocurrencies, add or remove cryptocurrencies from a tracking list, and export market data to a JSON file.

## Features
- **Track Cryptocurrencies:** Monitor real-time market data for selected cryptocurrencies.
- **Add New Cryptocurrency:** Include a new cryptocurrency in the tracking list.
- **Remove Cryptocurrency:** Eliminate a cryptocurrency from the tracking list.
- **Export Market Data:** Save current market data of all tracked cryptocurrencies to a JSON file.
- **Interactive Menu:** User-friendly command-line interface.

## Installation
### Prerequisites
- Go 1.18 or later
- Active internet connection (for fetching cryptocurrency data)

### Clone the Repository
```
git clone https://github.com/FG420/crypto-tracker.git
cd crypto-tracker
```

### Install Dependencies
Install necessary Go packages by running:
```
go mod tidy
```

## Usage
### Running the Application
To launch the application, execute:
```
go run main.go
```

### Menu Options
Upon starting the application, you'll encounter a menu with the following options:
1. Follow the inputted cryptocurrency market
2. Insert a new cryptocurrency to follow
3. Delete the inputted cryptocurrency from the list
4. Export the current market data for all tracked cryptocurrencies to a JSON file
5. Exit the application

## Example Workflow
1. Add a new cryptocurrency (e.g., BTC)
2. Begin monitoring BTC in real-time
3. Press Backspace to stop real-time monitoring
4. Export current market data of all tracked cryptocurrencies to a JSON file

## Managing Cryptocurrencies
Cryptocurrencies' information is stored in a `coins.json` file, created automatically on the first run. Manage this list using options [2] or [3] from the menu.

## File Structure
```
.
├── main.go               # Entry point of the application
├── crypto/               # Package for handling API calls and data fetching
├── types/                # Package for defining custom data types
├── coins.json            # JSON file storing the list of tracked cryptocurrencies
└── crypto-stock.json     # (Optional) JSON file containing the exported market data
```

## Contributing
Contributions are encouraged! Please feel free to open issues or submit pull requests to enhance the project.

### How to Contribute
1. Fork the repository
2. Create a new branch (`git checkout -b feature-branch`)
3. Commit your changes (`git commit -am 'Add new feature'`)
4. Push to the branch (`git push origin feature-branch`)
5. Create a new Pull Request

## License
This project is licensed under the MIT License - see the `LICENSE` file for details.
