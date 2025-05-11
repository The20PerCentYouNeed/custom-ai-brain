import { useState, useEffect } from "react";

export default function Documents() {
  const [documents, setDocuments] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  // Fetch documents on component mount
  useEffect(() => {
    const fetchDocuments = async () => {
      try {
        const response = await fetch("http://localhost:8080/documents");

        if (!response.ok) {
          throw new Error("Failed to fetch documents");
        }
        const data = await response.json();
        console.log(data);
        setDocuments(data);
      } catch (err) {
        setError(err.message);
      } finally {
        setLoading(false);
      }
    };

    fetchDocuments();
  }, []);

  if (loading) {
    return <p className="text-gray-500">Loading documents...</p>;
  }

  if (error) {
    return <p className="text-red-500">Error: {error}</p>;
  }



  return (
    <div>
      <h1 className="text-2xl font-bold mb-4">Documents</h1>
      {documents.length === 0 ? (
        <p className="text-gray-500">No documents found.</p>
      ) : (
        <table className="min-w-full bg-white shadow rounded overflow-hidden">
          <thead className="bg-gray-200 text-gray-700">
            <tr>
              <th className="text-left p-2">ID</th>
              <th className="text-left p-2">Title</th>
              <th className="text-left p-2">File</th>
              <th className="text-left p-2">Created At</th>
            </tr>
          </thead>
          <tbody>
            {documents.map((doc) => (
              <tr key={doc.id} className="border-t">
                <td className="p-2">{doc.id}</td>
                <td className="p-2">{doc.title}</td>
                <td className="p-2">{doc.file.uri}</td>
                <td className="p-2">{new Date(doc.created_at).toLocaleString()}</td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
}
