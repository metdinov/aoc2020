(ns aoc2020.day-01)

(defn part-a
  "Day 01 Part a"
  ([lines seen]
   (let [line (first lines)
         num (Integer/parseInt line)
         comp (- 2020 num)]
     (if (contains? seen comp) (* num comp)
         (recur (rest lines) (conj seen num)))))

  ([lines]
   (part-a lines #{})))

(defn contains-complement
  [seen num]
  (fn [el]
    (if (= el (- 2020 el num))
      (false)
      (contains? seen (- 2020 el num)))))

(defn find-first
  [pred coll]
  (first (filter pred coll)))

(defn find-sum-triplet
  [seen num]
  (find-first (contains-complement seen num) seen))

(defn part-b
  "Day 01 Part b"
  ([lines seen]
   (let [line (first lines)
         num (Integer/parseInt line)
         found (find-sum-triplet seen num)]
     (if found
       (* found num (- 2020 found num))
       (recur (rest lines) (conj seen num)))))

  ([lines]
   (part-b lines #{})))

