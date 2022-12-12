import { ReactComponent as MenuIcon } from "./assets/menu_1.svg";
import { ReactComponent as LogoutIcon } from "./assets/logout_1.svg";

interface IHeaderProps {
  sideBarStatus: boolean;
  setSideBarStatus: React.Dispatch<boolean>;
}

export default function Header({ sideBarStatus, setSideBarStatus }: IHeaderProps) {
  return (
    <header className="w-full min-w-fit shadow-sm flex justify-between items-center px-4 md:px-6 gap-4">
      <button
        className="transition-colors hover:bg-zinc-300/60 
        active:bg-zinc-300 focus:bg-zinc-300 focus:outline-none rounded-full p-2 my-2"
        onClick={() => setSideBarStatus(!sideBarStatus)}
      >
        <MenuIcon />
      </button>
      <button
        className="transition-colors hover:bg-zinc-300/60 
        active:bg-zinc-300 focus:bg-zinc-300 focus:outline-none rounded-full p-2"
      >
        <LogoutIcon />
      </button>
    </header>
  );
}
