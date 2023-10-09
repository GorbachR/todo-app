import { useRef, useState } from "react";
import Circle from "../Circle/Circle";
import TextareaAutosize from "react-textarea-autosize";
import { ReactComponent as Crossout } from "./assets/icon-cross.svg";
import { useMutation, useQuery, useQueryClient } from "react-query";
import { todoApi } from "../../App.tsx";

export type Todo = {
  id?: number;
  note: string;
  completed: boolean;
};

export type PutTodoParams = {
  todo: Todo;
  todoId: number;
};

type Response = {
  status: string;
  code?: number;
  message?: string;
  data?: Todo[];
};

export default function TodoList() {
  const [editTodo, setEditTodo] = useState<Todo | null | undefined>();
  const queryClient = useQueryClient();

  console.log(editTodo);

  const getTodoQuery = useQuery<Response>({
    queryKey: ["todos"],
    queryFn: async function (): Promise<Response> {
      const response = await todoApi.get<Response>("/todos");
      return response.data;
    },
  });

  const changeMutation = useMutation<Response, Error, Todo>({
    mutationFn: async function (todo) {
      const response = await todoApi.put(`/todos/${todo.id}`, todo);
      return response.data;
    },
    onSuccess: async () => {
      await queryClient.invalidateQueries({ queryKey: ["todos"] });
      setEditTodo(null);
    },
  });

  const deleteMutation = useMutation<Response, Error, number>({
    mutationFn: async function (todoId) {
      const response = await todoApi.delete(`/todos/${todoId}`);
      return response.data;
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["todos"] });
    },
  });

  function startEdit(todo: Todo) {
    setEditTodo(() => todo);
  }

  function handleDeleteClick(todoId: number) {
    deleteMutation.mutate(todoId);
  }

  function handleChangeClick(todoId: PutTodoParams) {
    changeMutation.mutate(todoId);
  }

  function handleCompletedChange(todo: Todo) {
    const newTodo = { ...todo, completed: !todo.completed };
    changeMutation.mutate(newTodo);
  }

  // if (isLoading) return <div className="">Loading...</div>;
  // if (error) return <div>An error has occurred:</div>;

  return (
    <ul className="rounded border border-transparent bg-todo-bg2 shadow-2xl">
      {getTodoQuery.data?.data?.map((todo: Todo) => {
        const editActive = !!(editTodo && todo.id === editTodo.id);
        return (
          <li
            key={todo.id}
            data-key={todo.id}
            className="group flex items-center gap-5 border-b border-b-todo-outline px-5"
          >
            <Circle
              id={`note-${todo.id}-check`}
              label={`Mark note ${todo.note} as done`}
              state={todo.completed}
              onChangeCallback={() => {
                handleCompletedChange(todo);
              }}
            />

            <div className="relative flex-grow">
              <TextareaAutosize
                className="block w-full resize-none truncate whitespace-pre-line bg-transparent py-5 focus:outline-none"
                style={todo.completed ? { textDecoration: "line-through" } : {}}
                value={editActive ? editTodo.note : todo.note}
                disabled={!editActive}
                maxRows={editActive ? undefined : 1}
                onBlur={() => {
                  if (editTodo) changeMutation.mutate(editTodo);
                }}
                onChange={(e) => {
                  setEditTodo((prev) => {
                    return { ...prev!, note: e.target.value };
                  });
                }}
                onKeyDown={(e) => {
                  switch (e.key) {
                    case "Enter":
                      if (!e.shiftKey && editTodo) {
                        changeMutation.mutate(editTodo);
                        e.preventDefault();
                      }
                      break;
                    case "Escape":
                      setEditTodo(null);
                      break;
                  }
                }}
              />
              {editTodo && todo.id === editTodo.id ? null : (
                <span
                  className="absolute bottom-0 left-0 right-0 top-0"
                  onClick={() => {
                    if (!todo.completed) startEdit(todo);
                  }}
                ></span>
              )}
            </div>

            <Crossout
              data-key={todo.id}
              className="crossout flex-shrink-0 opacity-0 group-hover:opacity-100"
              onClick={() => {
                const id = todo.id ? todo.id : 0;
                handleDeleteClick(id);
              }}
            />
          </li>
        );
      })}

      <li
        className="flex justify-between gap-5 p-5 [&_span]:text-xs
          [&_span]:text-todo-text-crossout"
      >
        <span>
          {getTodoQuery.data?.data
            ? getTodoQuery.data.data.length
            : 0 + " items left"}
        </span>
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
