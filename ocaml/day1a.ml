open Base
open Stdio

let contains set el = Set.exists set ~f:(fun x -> x = el)

let rec find_sum_pair ?(set = Set.empty (module Int)) input_ch =
  let line = In_channel.input_line input_ch in
  match line with
  | None -> failwith "Not found"
  | Some x ->
    let num = Int.of_string x in
    if contains set num then num * (2020 - num)
    else find_sum_pair ~set:(Set.add set (2020 - num)) input_ch

let solve filename = In_channel.with_file filename ~f:find_sum_pair

let () = "../data/day1a.txt" |> solve |> printf "%d\n"
