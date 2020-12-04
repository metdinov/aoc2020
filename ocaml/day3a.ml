open Base
open Stdio

let is_tree line pos =
  let c = line.[pos] in
  Char.(c = '#')

let rec count_trees ?(total = 0) ?(pos = 0) input_ch =
  let line = In_channel.input_line input_ch in
  match line with
  | None -> total
  | Some l ->
      let length = String.length l in
      let new_pos = (pos + 3) % length in
      if is_tree l pos then count_trees ~total:(total + 1) ~pos:new_pos input_ch
      else count_trees ~total ~pos:new_pos input_ch

let solve filename = In_channel.with_file filename ~f:count_trees

let () = "../data/day3a.txt" |> solve |> printf "%d\n"
