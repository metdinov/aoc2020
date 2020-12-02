(ns aoc2020.core
  (:require [aoc2020.day-01]
            [aoc2020.day-02]
            [aoc2020.day-03]
            [aoc2020.day-04]
            [aoc2020.day-05]
            [aoc2020.day-06]
            [aoc2020.day-07]
            [aoc2020.day-08]
            [aoc2020.day-09]
            [aoc2020.day-10]
            [aoc2020.day-11]
            [aoc2020.day-12]
            [aoc2020.day-13]
            [aoc2020.day-14]
            [aoc2020.day-15]
            [aoc2020.day-16]
            [aoc2020.day-17]
            [aoc2020.day-18]
            [aoc2020.day-19]
            [aoc2020.day-20]
            [aoc2020.day-21]
            [aoc2020.day-22]
            [aoc2020.day-23]
            [aoc2020.day-24]
            [aoc2020.day-25]))

(defn read-input
  [day]
  (-> day clojure.java.io/resource slurp clojure.string/split-lines))

(defn -main
  "Used to dispatch tasks from the command line.
  
  lein run day1a"
  [part]
  (case part
    "day1a" (-> "day1a.txt" read-input aoc2020.day-01/part-a println)
    "day1b" (-> "day1b.txt" read-input aoc2020.day-01/part-b println)
    (println "not found")))

