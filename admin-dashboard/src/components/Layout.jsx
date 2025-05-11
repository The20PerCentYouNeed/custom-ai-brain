import Navigation from "./Navigation";
import { Outlet } from 'react-router-dom';

export default function Layout() {
  return (
    <div className="flex h-screen">
      <Navigation />
      <main className="flex-1 bg-gray-100 p-6 overflow-auto mt-16 md:mt-0 w-full">
        <Outlet />
      </main>
    </div>
  );
}
