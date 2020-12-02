open Base
open Stdio

let extract s =
  Caml.Scanf.sscanf s "%d-%d %c: %s" (fun min max c passw ->
      (min, max, c, passw))

let char_count char str = String.count ~f:(fun c -> Char.(char = c)) str

let rec count_valid_passwords ?(valid = 0) input_ch =
  let line = In_channel.input_line input_ch in
  match line with
  | None -> valid
  | Some l ->
    let min, max, c, passwd = extract l in
    let count = char_count c passwd in
    if count >= min && count <= max then
      count_valid_passwords ~valid:(valid + 1) input_ch
    else count_valid_passwords ~valid input_ch

let solve filename = In_channel.with_file filename ~f:count_valid_passwords

let () = "../data/day2a.txt" |> solve |> printf "%d\n"
