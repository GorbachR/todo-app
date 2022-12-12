import useGetNotes from "../../hooks/useGetNotes";

interface ISidebar {
  sideBarStatus: boolean;
}

export default function Sidebar({ sideBarStatus }: ISidebar) {
  const visibleStyle = "w-80 opacity-100 p-4 flex-none";
  const collapsedStyle = "w-0 opacity-0";
  const { data: notesData } = useGetNotes();

  console.log(notesData);

  return (
    <div
      className={`${
        sideBarStatus ? visibleStyle : collapsedStyle
      } transition-all py-4 shadow-md bg-neutral-100 overflow-hidden mx-auto whitespace-nowrap`}
    >
      <h1 className="">
        <span className="text-purple-500">Notes</span> App
      </h1>
      <ul>
        <li>firstNote</li>
        <li>secondNote</li>
        <li>firstNote</li>
        <li>secondNote</li>
        <li>firstNote</li>
        <li>secondNote</li>
        <li>firstNote</li>
        <li>secondNote</li>
      </ul>
    </div>
  );
}
