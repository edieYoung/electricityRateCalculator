# electricityRateCalculator
This helps you to calculate the eletricity cost of every single month, cost in total(counting possible termination fee) and the average cost according to the rate plan you can define, and it is based on the original/estimated usate report.

## Define the Rate Plan
You can calculate 
Modify the main function in calculator.go
```
// instantiate a rate plan {rate@500, rate@1000, terminationFee, leasingLength}
rp := RatePlan{0.13, 0.12, 175, 6}
// calculate the cost according to an original/estimated usage report (csv file)
costMap := rp.Calculate("usage_2019-02-20.csv") 

```

## Sample Usage Report Format
 the sample report is having the start date(yy-mm-dd) in column 2 and usage(number) located in column 4.

 ## Run it and Get Results
Simply build it and run in terminal: 
```
    $go build
    $./eletricityRateCalculator
```
you will see the results like:
```
StartDate Value (kWh)
map[01:74.88849 02:47.429592 08:43.422337 09:109.60759 10:80.17004 11:74.42132 12:67.05511 Average:70.999214 Total:496.99448]
```

which tells you about the cost information all in a map.