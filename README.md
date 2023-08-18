# Breaklist - A morning report generator for thermal printers
Breaklist is a set of tools for generating a morning report on a thermal printer. The morning report includes:

- A task list 
- Reminders
- Weather forecast
- Summary of latest top articles on Hacker News

The project also features a user-friendly web application that facilitates task management.

## Features

- Tasks are stored in a `tasks.list` file in plain text format
- Reminders are stored in a `reminders.list` file using crontab format
- The weather forecast is retrieved from [Tomorrow.io](https://docs.tomorrow.io/reference/welcome)'s API (requires API key)
- Hacker News summaries are powered by [polyrabbit's hacker-news-digest](https://github.com/polyrabbit/hacker-news-digest)
- A web app provides an interface for managing tasks
- The morning report is generated as a PDF using [wkhtmltopdf](https://github.com/wkhtmltopdf/wkhtmltopdf)

## Getting started

### Prerequisites

- [wkhtmltopdf](https://wkhtmltopdf.org/downloads.html) needs to be installed. It can be downloaded from [here](https://wkhtmltopdf.org/downloads.html).
- A Tomorrow.io API key

### Installation
To use Breaklist, you can either download the latest release from the GitHub releases page or build it from source:
### Downloding latest release
Download the most recent release from the following link:
https://github.com/alibahmanyar/breaklist/releases/latest

### Building from source
```sh
git clone git@github.com:alibahmanyar/breaklist.git
cd breaklist
make setup
make
```
The compiled binaries will be available in the `build` directory.

### Usage
Before running Breaklist, duplicate the `.env.example` file and rename it to `.env`, then populate the variables accordingly.

Once the variables are set, Breaklist can be operated using the following commands:

- The `reportGenerator` executable will generate the reporrt as `breaklist.pdf` which can be printed using a thermal printer.

- The `webserver` executable will run the web application to manage the tasks. The web app will be available at `:3030`.

### tasks.list and reminders.list
The `tasks.list` file contains the list of current tasks in plain text format, with each task on a separate line.

The `reminders.list` file holds the reminders in a crontab-style format, with each reminder on a separate line.
Example:
```
#.---------- day of month (1 - 31)
#|  .------- month (1 - 12) OR jan,feb,mar,apr ...
#|  |  .---- day of week (0 - 6) (Sunday=0 or 7) OR sun,mon,tue,wed,thu,fri,sat
#|  |  |

* * *|Reminder 1 (Every day)
*/2 * *|Reminder 2 (Every other day)
* * 6,0|Reminder 2 (On Saturdays and Sundays)
```
