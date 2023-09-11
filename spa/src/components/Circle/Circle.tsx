type CircleProps = {
  checkPossible: boolean;
  checkState?: boolean;
  id: string;
  label: string;
};

export default function Circle({
  checkPossible = false,
  checkState = false,
  id,
  label,
}: CircleProps) {
  return (
    <div className="relative h-[1.5rem] w-[1.5rem] flex-shrink-0">
      <label htmlFor={id} className="sr-only">
        {label}
      </label>
      <input
        id={id}
        type="checkbox"
        className="peer absolute z-10 h-full w-full cursor-pointer appearance-none rounded-full bg-none bg-cover
        bg-center bg-no-repeat checked:bg-[url('./components/Circle/assets/icon-check.svg')]"
        style={{ backgroundSize: "0.75rem 0.5625em" }}
        aria-checked={checkState}
        defaultChecked={checkState}
        disabled={checkPossible}
        onChange={() => {
          return;
        }}
      />
      <span
        className="absolute h-full w-full rounded-full bg-todo-outline before:absolute before:left-1/2 
      before:top-1/2 before:h-[1.375rem] before:w-[1.375rem] before:-translate-x-1/2 before:-translate-y-1/2 
      before:rounded-full before:bg-todo-bg2 peer-checked:bg-gradient-to-br peer-checked:from-todo-circle-hover1 
      peer-checked:to-todo-circle-hover2 peer-checked:before:bg-gradient-to-br peer-checked:before:from-todo-circle-hover1 
      peer-checked:before:to-todo-circle-hover2 peer-hover:bg-gradient-to-br peer-hover:from-todo-circle-hover1 
      peer-hover:to-todo-circle-hover2"></span>
    </div>
  );
}
