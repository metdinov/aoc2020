open Base
open Stdio

let extract s =
  Caml.Scanf.sscanf s "%d-%d %c: %s" (fun min max c passw ->
      (min, max, c, passw))

let is_valid str char (pos1, pos2) =
  let c1 = str.[pos1 - 1] in
  let c2 = str.[pos2 - 1] in
  Poly.(c1 = char <> (c2 = char))

let rec count_valid_passwords ?(valid = 0) input_ch =
  let line = In_channel.input_line input_ch in
  match line with
  | None -> valid
  | Some l ->
    let pos1, pos2, c, passwd = extract l in
    if is_valid passwd c (pos1, pos2) then
      count_valid_passwords ~valid:(valid + 1) input_ch
    else count_valid_passwords ~valid input_ch

let solve filename = In_channel.with_file filename ~f:count_valid_passwords

let () = "../data/day2b.txt" |> solve |> printf "%d\n"
