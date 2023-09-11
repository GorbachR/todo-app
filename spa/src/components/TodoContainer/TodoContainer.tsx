import { useContext, useState } from "react";
import { ReactComponent as Moon } from "./assets/icon-moon.svg";
import { ReactComponent as Sun } from "./assets/icon-sun.svg";
import ThemeContext from "../../context/ThemeContext";
import NotesList from "../NotesList/NotesList";
import TextareaAutosize from "react-textarea-autosize";
import Circle from "../Circle/Circle";

export default function TodoContainer() {
  const themeState = useContext(ThemeContext);
  const [input, setInput] = useState<string>("test");

  return (
    <div className="mx-auto flex w-[33%] min-w-[33.75rem] max-w-[33.75rem] flex-col px-8 pt-[4.75rem]">
      <div className="mb-12 flex justify-between">
        <h1>TODO</h1>
        <div
          className="cursor-pointer"
          onClick={(e) => {
            themeState?.setTheme((theme) =>
              theme === "light" ? "dark" : "light"
            );
          }}>
          {themeState?.theme === "light" ? <Moon /> : <Sun />}
        </div>
      </div>
      <div className="mb-6 flex items-center gap-5 rounded border border-transparent bg-todo-bg2 p-5">
        <Circle checkPossible={false} id="save-note" label="Save todo note" />
        <TextareaAutosize
          className="flex-grow resize-none bg-transparent focus:outline-none"
          placeholder="Create a new todo..."
          autoFocus
          value={input}
          onChange={(e) => setInput(e.target.value)}
        />
      </div>
      <NotesList />
    </div>
  );
}
