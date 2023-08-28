 ## Golang CLI Reminder

This is a simple command line reminder application written in Go. It allows you to set a reminder for a specific time and date, and it will display a notification when the time comes.

### Prerequisites

To use this application, you will need to have Go installed on your system. You can download Go from the [official website](https://golang.org/dl/).

Once you have Go installed, you can clone this repository by running the following command:

```
git clone https://github.com/example/golang-cli-reminder.git
```

### Usage

To use the application, you will need to provide the time and date of the reminder, as well as the message you want to be displayed. You can do this by running the following command:

```
main.go <hh:mm> <text-message>
```

For example, to set a reminder for 10:00 AM on January 1, 2023, with the message "Take out the trash", you would run the following command:

```
main.go 10:00 Take out the trash
```

The application will then display a notification at 10:00 AM on January 1, 2023, with the message "Take out the trash".