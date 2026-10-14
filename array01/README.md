# Telemetry Temperature Analysis (Array Lab 01)

A concise Go program analyzing a fixed-size telemetry batch (`[6]int`) to determine minimum and maximum temperatures without dynamic memory allocations.

## Problem

Given a batch of 6 temperature readings:
- Find the **minimum** temperature.
- Find the **maximum** temperature.

## Example

- **Input:** `[6]int{40, 15, -24, 0, 4, 6}`
- **Output:** `min: -24 , max: 40`

## Usage

```bash
go run main.go