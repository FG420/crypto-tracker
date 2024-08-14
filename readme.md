Crypto Tracker
Overview
Crypto Tracker is a command-line application written in Go that allows users to manage and track the prices of various cryptocurrencies. The application provides options to view the current market data of specific cryptocurrencies, add or remove cryptocurrencies from a list, and export the current market data to a JSON file.

Features
Track Cryptocurrencies: Follow real-time market data for selected cryptocurrencies.
Add New Cryptocurrency: Add a new cryptocurrency to the tracking list.
Remove Cryptocurrency: Remove a cryptocurrency from the tracking list.
Export Market Data: Save the current market data of all tracked cryptocurrencies to a JSON file.
Interactive Menu: Easy-to-use command-line interface.
Installation
Prerequisites
Go 1.18 or later
A working internet connection (for fetching cryptocurrency data)
Clone the Repository
CopyReplit
git clone https://github.com/FG420/crypto-tracker.git
cd crypto-tracker
Install Dependencies
This project relies on a few Go packages. You can install them by running:

CopyReplit
go mod tidy
Usage
Running the Application
To start the application, run:

CopyReplit
go run main.go
Menu Options
After starting the application, you'll be presented with a menu:

Follow the inputted cryptocurrency market: Select this option to continuously monitor the market data for a specific cryptocurrency.
Insert a new cryptocurrency to follow: Add a new cryptocurrency to the tracking list.
Delete the inputted cryptocurrency from the list: Remove a cryptocurrency from the tracking list.
Export the current market data for all tracked cryptocurrencies to a JSON file.
Exit: Close the application.
Example Workflow
Select option [2] to add a new cryptocurrency to the list (e.g., BTC).
Select option [1] to start following BTC in real-time.
Press Backspace to stop following the market data.
Select option [4] to export the current market data of all tracked cryptocurrencies to a JSON file.
Managing Cryptocurrencies
Cryptocurrencies are stored in a coins.json file, which is automatically created when the application is first run. To view or manage this list, use options [2] or [3] from the menu.

File Structure
CopyReplit
.
├── main.go               # Entry point of the application
├── crypto/               # Package for handling API calls and data fetching
├── types/                # Package for defining custom data types
├── coins.json            # JSON file storing the list of tracked cryptocurrencies
└── crypto-stock.json     # (Optional) JSON file containing the exported market data
Contributing
Contributions are welcome! Feel free to open issues or submit pull requests to improve the project.

How to Contribute
Fork the repository.
Create a new branch (git checkout -b feature-branch).
Commit your changes (git commit -am 'Add new feature').
Push to the branch (git push origin feature-branch).
Create a new Pull Request.
License
This project is licensed under the MIT License - see the LICENSE file for details.
