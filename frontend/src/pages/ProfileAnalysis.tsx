import { useState } from 'react';
import { FileUpload } from '../components/FileUpload';
import { profileApi } from '../api/client';

export const ProfileAnalysis = () => {
  const [file, setFile] = useState<File | null>(null);
  const [loading, setLoading] = useState(false);
  const [result, setResult] = useState<string | null>(null);

  const handleAnalyze = async () => {
    if (!file) return;

    setLoading(true);
    try {
      const html = await profileApi.analyzeSimpleProfile(file);
      setResult(html);
    } catch (error) {
      console.error('Analysis failed:', error);
      alert('Analysis failed. Please check the console for details.');
    } finally {
      setLoading(false);
    }
  };

  if (result) {
    return (
      <div className="w-full h-screen">
        <div className="mb-4">
          <button
            onClick={() => setResult(null)}
            className="px-4 py-2 bg-primary text-white rounded-lg hover:bg-primary-dark"
          >
            ← Back to Upload
          </button>
        </div>
        <iframe
          srcDoc={result}
          className="w-full h-full border-0"
          title="Analysis Result"
        />
      </div>
    );
  }

  return (
    <div className="max-w-4xl">
      <h1 className="text-3xl font-bold text-gray-900 mb-4">
        Simple Profile Analysis
      </h1>

      <div className="bg-white rounded-lg shadow-md p-6 mb-6">
        <h2 className="text-xl font-semibold mb-2">Tool Purpose</h2>
        <p className="text-gray-700">
          Searchable, sortable, exportable view of all operators in a profile.json file.
        </p>
      </div>

      <div className="bg-white rounded-lg shadow-md p-6">
        <h2 className="text-xl font-semibold mb-4">Upload Form</h2>

        <div className="space-y-4">
          <FileUpload
            label="Profile File"
            accept=".tar,.gz,.tgz,.zip,.json"
            onChange={(files) => setFile(files[0] || null)}
            helperText="Attach one profile in a text file or archive (.tar, .gz, .tgz, .zip, .json)"
          />

          <button
            onClick={handleAnalyze}
            disabled={!file || loading}
            className="w-full px-6 py-3 bg-primary text-white font-medium rounded-lg hover:bg-primary-dark disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
          >
            {loading ? 'Analyzing...' : 'Analyze'}
          </button>
        </div>
      </div>
    </div>
  );
};
