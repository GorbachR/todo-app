import { MouseEventHandler, useRef, useState } from "react";
import Circle from "../Circle/Circle";
import TextareaAutosize from "react-textarea-autosize";
import { ReactComponent as Crossout } from "./assets/icon-cross.svg";
import { useMutation, useQuery } from "react-query";
import axios from "axios";
const apiUrl = import.meta.env.BASE_URL;

type Note = {
  id: string;
  content: string;
  active: boolean;
};

export default function NotesList() {
  const [editNote, setEditNote] = useState<string | null | undefined>(null);

  const { isLoading, error, data } = useQuery({
    queryKey: ["todo-notes"],
    queryFn: fetchNotes,
  });

  const deleteMutation = useMutation({
    mutationFn: (note) => {
      return axios.delete(`${apiUrl}/notes/${note}`);
    },
  });

  function fetchNotes(): Promise<Note[]> {
    return axios.get(`${apiUrl}/notes`).then((res) => res.data);
  }

  function handleDoubleClick(e: React.MouseEvent<HTMLDivElement>) {
    setEditNote(
      (e.target as HTMLElement).closest("[data-key]")?.getAttribute("data-key")
    );
  }

  function handleCrossoutClick(e: React.MouseEvent<SVGElement>) {
    console.log(e.target);
  }

  // if (isLoading) return <div className="">Loading...</div>;
  // if (error) return <div>An error has occurred:</div>;

  return (
    <ul className="rounded border border-transparent bg-todo-bg2 shadow-2xl">
      {data?.map((input) => (
        <li
          key={input.id}
          data-key={input.id}
          className="group flex items-center gap-5 border-b border-b-todo-outline p-5">
          <Circle
            checkPossible={false}
            checkState={input.active}
            id={`note-${input.id}-check`}
            label={`Mark note ${input.content} as done`}
          />

          <div className="flex-grow overflow-hidden">
            {input.id === editNote ? (
              <TextareaAutosize
                className="block resize-none bg-transparent focus:outline-none"
                value={input.content}
                autoFocus
                onFocus={(e) => e.currentTarget.select()}
              />
            ) : (
              <p
                className="flex-grow truncate"
                onDoubleClick={handleDoubleClick}>
                {input.content}
              </p>
            )}
          </div>

          <Crossout
            data-key={input.id}
            className="crossout flex-shrink-0 opacity-0 group-hover:opacity-100"
            onClick={handleCrossoutClick}
          />
        </li>
      ))}

      <li
        className="flex justify-between gap-5 p-5 [&_span]:text-xs 
          [&_span]:text-todo-text-crossout">
        <span>{data ? data.length : 0 + " items left"}</span>
        <div className="flex justify-between gap-4 font-bold">
          <span>All</span>
          <span>Active</span>
          <span>Completed</span>
        </div>
        <span>Clear completed</span>
      </li>
    </ul>
  );
}
