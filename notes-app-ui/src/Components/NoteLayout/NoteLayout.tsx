import { useRef, useState } from "react";
import { Outlet } from "react-router-dom";
import Footer from "../Footer/Footer";
import Header from "../Header/Header";
import Sidebar from "../Sidebar/Sidebar";

export default function NoteLayout() {
  const [sideBarStatus, setSideBarStatus] = useState<boolean>(true);

  return (
    <div className="min-h-screen h-1 flex">
      <Sidebar sideBarStatus={sideBarStatus} />
      <div className="flex flex-col flex-auto">
        <Header sideBarStatus={sideBarStatus} setSideBarStatus={setSideBarStatus} />
        <Outlet />
        <Footer />
      </div>
    </div>
  );
}
