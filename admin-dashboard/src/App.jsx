import { Routes, Route } from 'react-router-dom';
import Layout from './components/Layout';
import Home from './pages/Home';
import UploadFiles from './pages/UploadFiles';
import KnowledgeBase from './pages/KnowledgeBase';
import Chat from './pages/Chat';

export default function App() {
  return (
    <Routes>
      <Route path="/" element={<Layout />}>
        <Route index element={<Home />} />
        <Route path="upload" element={<UploadFiles />} />
        <Route path="knowledge" element={<KnowledgeBase />} />
        <Route path="chat" element={<Chat />} />
      </Route>
    </Routes>
  );
}
