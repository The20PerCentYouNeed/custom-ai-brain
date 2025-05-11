import { useState } from "react";
import { Home, Book, X, Menu, MessageCircle, Folder } from "lucide-react";
import { Link } from "react-router-dom";

export default function Navigation() {
  const [isSidebarOpen, setIsSidebarOpen] = useState(false);
  const toggleSidebar = () => setIsSidebarOpen(!isSidebarOpen);

  const [isDropdownOpen, setIsDropdownOpen] = useState(false);
  const openDropdown = () => setIsDropdownOpen(true);
  const closeDropdown = () => setIsDropdownOpen(false);

  const SidebarContent = () => (
    <div className="w-64 bg-gray-900 text-white flex flex-col h-full">
      <div className="p-4 text-2xl font-bold border-b border-gray-700 flex justify-between items-center">
        Admin
        <button
          className="md:hidden"
          onClick={toggleSidebar}
          aria-label="Close menu"
        >
          <X size={24} />
        </button>
      </div>
      <nav className="flex-1 p-4 space-y-4">
        <Link
          to="/"
          onClick={closeDropdown}
          className="flex items-center gap-2 hover:text-blue-400"
        >
          <Home size={20} /> Home
        </Link>

        <div className="block">
          <button
            onClick={openDropdown}
            className="flex items-center justify-between w-full hover:text-blue-400"
          >
            <Link to="/documents" className="flex items-center gap-2 hover:text-blue-400">
                <Folder size={20} /> Documents
            </Link>
          </button>

          {isDropdownOpen && (
            <div className="ml-6 mt-3 space-y-2 text-sm">
              <Link to="/documents/chunks" className="block hover:text-blue-400">
                Chunks
              </Link>
              <Link to="/documents/files" className="block hover:text-blue-400">
                Files
              </Link>
            </div>
          )}
        </div>

        <Link
          to="/chat"
          onClick={closeDropdown}
          className="flex items-center gap-2 hover:text-blue-400"
        >
          <MessageCircle size={20} /> Chat With AI
        </Link>
        <Link
          to="/knowledge"
          onClick={closeDropdown}
          className="flex items-center gap-2 hover:text-blue-400"
        >
          <Book size={20} /> Knowledge Base
        </Link>
      </nav>
    </div>
  );

  return (
    <>
      <div className="md:hidden fixed top-0 left-0 right-0 bg-gray-900 text-white flex items-center justify-between p-4 z-20">
        <span className="text-xl font-bold">Admin</span>
        <button onClick={toggleSidebar} aria-label="Open menu">
            <Menu size={24} />
        </button>
      </div>

      <aside className="hidden md:block">{<SidebarContent />}</aside>

      {isSidebarOpen && (
        <div className="fixed inset-0 z-30 flex">
          <div
            className="absolute inset-0 bg-black opacity-50"
            onClick={toggleSidebar}
          />
          <aside className="relative z-40">{<SidebarContent />}</aside>
        </div>
      )}
    </>
  );
}
