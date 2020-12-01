open Base
open Stdio

let parse_int input_ch = 
  let line = In_channel.input_line input_ch in
  match line with
  | None -> failwith "Not found"
  | Some x -> Int.of_string x

let contains set el = Set.exists set ~f:(fun x -> x = el)

let contains_complement set num el = (el <> 2020 - num) && contains set (2020 - num - el)

let check_sum set num = Set.find set ~f:(contains_complement set num)

let rec find_sum_triplet ?seen:(seen=Set.empty (module Int)) input_ch =
  let num = parse_int input_ch in
  match check_sum seen num with
  | None -> find_sum_triplet ~seen:(Set.add seen num) input_ch
  | Some y -> y * num * (2020 - y - num)

let solve filename =
  In_channel.with_file filename ~f:find_sum_triplet

let () = 
  "../data/day1b.txt"
  |> solve
  |> printf "%d\n";
