import { useContext, useRef, useState } from "react";
import { ReactComponent as Moon } from "./assets/icon-moon.svg";
import { ReactComponent as Sun } from "./assets/icon-sun.svg";
import ThemeContext from "../../context/ThemeContext";
import TodoList from "../NotesList/TodoList.tsx";
import TextareaAutosize from "react-textarea-autosize";
import Circle from "../Circle/Circle";
import { useMutation, useQueryClient } from "react-query";
import { todoApi } from "../../App.tsx";
import { Todo } from "../NotesList/TodoList.tsx";

export default function TodoContainer() {
  const themeState = useContext(ThemeContext);
  const inputRef = useRef<HTMLTextAreaElement>(null);
  const [createTodo, setCreateTodo] = useState<boolean>(false);
  const queryClient = useQueryClient();

  const createMutation = useMutation<Response, Error, Todo>({
    mutationFn: async function (todo) {
      const response = await todoApi.post(`/todos`, todo);
      return response.data;
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["todos"] });
    },
  });

  function createTodoCallback() {
    setCreateTodo(true);
    createMutation.mutate({ note: inputRef.current!.value, completed: false });
    setTimeout(() => {
      setCreateTodo(false);
    }, 200);
    inputRef.current!.value = "";
  }

  return (
    <div className="mx-auto flex w-[33%] min-w-[33.75rem] max-w-[33.75rem] flex-col px-8 pt-[4.75rem]">
      <div className="mb-12 flex justify-between">
        <h1>TODO</h1>
        <div
          className="cursor-pointer"
          onClick={() => {
            themeState?.setTheme((theme) =>
              theme === "light" ? "dark" : "light"
            );
          }}
        >
          {themeState?.theme === "light" ? <Moon /> : <Sun />}
        </div>
      </div>
      <div className="mb-6 flex items-center gap-5 rounded border border-transparent bg-todo-bg2 p-5">
        <Circle
          id="save-note"
          label="Save todo note"
          state={createTodo}
          onChangeCallback={createTodoCallback}
          onMouseDownCallback={() => {
            setCreateTodo(true);
          }}
          onMouseLeaveCallback={() => {
            setCreateTodo(false);
          }}
        />
        <TextareaAutosize
          className="flex-grow resize-none bg-transparent focus:outline-none"
          placeholder="Create a new todo..."
          autoFocus
          ref={inputRef}
          onKeyDown={(e) => {
            switch (e.key) {
              case "Enter":
                if (!e.shiftKey) {
                  createTodoCallback();
                  e.preventDefault();
                }
                break;
            }
          }}
        />
      </div>
      <TodoList />
    </div>
  );
}
