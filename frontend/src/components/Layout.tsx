import { Link, Outlet } from 'react-router-dom';
import { useAbout } from '../hooks/useAbout';

export const Layout = () => {
  const { data: about } = useAbout();

  const navigation = [
    { name: 'Home', href: '/', icon: '🏠' },
    { name: 'Simple Profile', href: '/profile', icon: '📊' },
    { name: 'Detailed Profile', href: '/profile-detailed', icon: '📈' },
    { name: 'Profile Comparison', href: '/profiles-comparison', icon: '🔄' },
    { name: 'Queries.json', href: '/queries-json', icon: '📊' },
    { name: 'Schema Generation', href: '/schema', icon: '🗄️' },
    { name: 'IOStat', href: '/iostat', icon: '💻' },
    { name: 'Threaded Top', href: '/top', icon: '⬆️' },
  ];

  return (
    <div className="min-h-screen bg-gray-50">
      {/* Sidebar */}
      <div className="fixed inset-y-0 left-0 w-64 bg-primary text-white">
        <div className="flex flex-col h-full">
          <div className="p-6">
            <h1 className="text-2xl font-bold">DQD {about?.version}</h1>
            <p className="text-sm text-primary-light mt-1">Dremio Query Doctor</p>
          </div>

          <nav className="flex-1 px-4 space-y-1">
            {navigation.map((item) => (
              <Link
                key={item.name}
                to={item.href}
                className="flex items-center px-4 py-3 text-sm font-medium rounded-lg hover:bg-primary-dark transition-colors"
              >
                <span className="mr-3">{item.icon}</span>
                {item.name}
              </Link>
            ))}
          </nav>

          <div className="p-4 text-xs text-primary-light border-t border-primary-dark">
            <p>© 2024 Dremio</p>
            <a
              href="https://github.com/rsvihladremio/dqd"
              target="_blank"
              rel="noopener noreferrer"
              className="text-primary-light hover:text-white underline"
            >
              GitHub Repository
            </a>
          </div>
        </div>
      </div>

      {/* Main content */}
      <div className="ml-64">
        <main className="p-8">
          <Outlet />
        </main>
      </div>
    </div>
  );
};
