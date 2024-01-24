# Weather Forecast API

This repository contains a RESTful API built in Golang using GoFiber to fetch weather forecasts for all provinces in Indonesia from BMKG (Badan Meteorologi, Klimatologi, dan Geofisika).

## Table of Contents
- [Introduction](#introduction)
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Contribution](#contribution)
- [License](#license)

## Introduction

This API provides a simple and efficient way to access weather forecasts for all provinces in Indonesia. It utilizes the data provided by BMKG to deliver accurate and up-to-date information.

## Features

- Retrieve weather forecasts for all provinces in Indonesia.
- Data updated per 3 days.
- Lightweight and fast, built with GoFiber.

## Tech Stack

- Golang
- GoFiber
- Viper

## Getting Started

### Prerequisites

- Golang installed on your machine.

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/bayazidsustami/bmkg-api
   cd bmkg-api  
2. Install the dependencies
    ```go mod download```
3. Open file `/config/config.env` and adjust desired host and port
 
## Usage
1. Run The Application :
    ```
    go run main.go
    ```
2. Acces the api :
    `http://localhost:8000` (default port in config file)

## Api Endpoints
- GET `/api/weather/provinces`: Get weather forecasts for all provinces.
- GET `/api/weather/provinces/{provinceId}`: Get weather forecast for a specific province.

### List Province
| Province ID | Province |
| :------: | :------ |
| 1 | Aceh |
| 2 | Bali |
| 3 | Bangka Belitung |
| 4 | Banten |
| 5 | Bengkulu |
| 6 | DI Yogyakarta |
| 7 | DKI Jakarta |
| 8 | Gorontalo |
| 9 | Jambi |
| 10 | Jawa Barat |
| 11 | Jawa Tengah |
| 12 | JawaTimur |
| 13 | Kalimantan Barat |
| 14 | Kalimantan Selatan |
| 15 | Kalimantan Tengah |
| 16 | Kalimantan Timur |
| 17 | Kalimantan Utara |
| 18 | Kepulauan Riau |
| 19 | Lampung |
| 20 | Maluku |
| 21 | Maluku Utara |
| 22 | Nusa Tenggara Barat |
| 23 | Nusa Tenggara Timur |
| 24 | Papua |
| 25 | Papua Barat |
| 26 | Riau |
| 27 | Sulawesi Barat |
| 28 | Sulawesi Selatan |
| 29 | Sulawesi Tengah |
| 30 | Sulawesi Tenggara |
| 31 | Sulawesi Utara |
| 32 | Sumatera Barat |
| 33 | Sumatera Selatan |
| 34 | Sumatera Utara |

## Contribution
Feel free to contribute by opening issues or submitting pull requests. Contributions are always welcome!

## License
This project is licensed under the MIT License - see the LICENSE file for details.