import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { Layout } from './components/Layout';
import { Home } from './pages/Home';
import { ProfileAnalysis } from './pages/ProfileAnalysis';

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      refetchOnWindowFocus: false,
      retry: 1,
    },
  },
});

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Layout />}>
            <Route index element={<Home />} />
            <Route path="profile" element={<ProfileAnalysis />} />
            <Route path="profile-detailed" element={<div>Detailed Profile (Coming Soon)</div>} />
            <Route path="profiles-comparison" element={<div>Profile Comparison (Coming Soon)</div>} />
            <Route path="queries-json" element={<div>Queries.json (Coming Soon)</div>} />
            <Route path="schema" element={<div>Schema Generation (Coming Soon)</div>} />
            <Route path="iostat" element={<div>IOStat (Coming Soon)</div>} />
            <Route path="top" element={<div>Threaded Top (Coming Soon)</div>} />
          </Route>
        </Routes>
      </BrowserRouter>
    </QueryClientProvider>
  );
}

export default App;
