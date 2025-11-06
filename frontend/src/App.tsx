import { useEffect, useState } from 'react';
import { Loader } from './components/Loader';
import { SectionRenderer } from './components/SectionRenderer';
import { SECTIONS } from './sections';

export function App() {
  const [activeSection, setActiveSection] = useState(SECTIONS[0]?.id ?? 'home');
  const [version, setVersion] = useState('');
  const [loading, setLoading] = useState(false);
  const [sidebarOpen, setSidebarOpen] = useState(false);

  useEffect(() => {
    let cancelled = false;

    async function fetchVersion() {
        try {
          const response = await fetch('/about.json');
          if (!response.ok) {
            throw new Error(`HTTP ${response.status}`);
          }
          const data = await response.json();
          if (!cancelled) {
            setVersion(data?.version ? `DQD ${data.version}` : 'DQD');
          }
        } catch (error) {
          if (!cancelled) {
            console.warn('Unable to fetch version information', error);
          }
        }
    }

    fetchVersion();
    return () => {
      cancelled = true;
    };
  }, []);

  const handleNav = (sectionId: string) => {
    setActiveSection(sectionId);
    setSidebarOpen(false);
  };

  const handleFormSubmit = () => {
    setLoading(true);
    setSidebarOpen(false);
  };

  const currentSection = SECTIONS.find((section) => section.id === activeSection) ?? SECTIONS[0];

  return (
    <div className="app-shell">
      <aside className={`sidebar ${sidebarOpen ? 'sidebar--open' : ''}`} aria-label="Main navigation">
        <div className="sidebar__header">
          <span className="sidebar__product">Dremio Query Doctor</span>
          <span className="sidebar__tagline">Guided diagnostics for your support files</span>
          <span className="sidebar__version">{version || 'unknown version'}</span>
        </div>
        <nav className="sidebar__nav">
          {SECTIONS.map((section) => (
            <button
              key={section.id}
              type="button"
              className={`sidebar__link ${activeSection === section.id ? 'is-active' : ''}`}
              onClick={() => handleNav(section.id)}
            >
              <span className="material-symbols-rounded">{section.navIcon}</span>
              <span>{section.navLabel}</span>
            </button>
          ))}
        </nav>
        <div className="sidebar__footer">
          <div>Need additional help?</div>
          <a href="https://github.com/rsvihladremio/dqd" target="_blank" rel="noreferrer">
            Contribute on GitHub
          </a>
        </div>
      </aside>

      <main className="main">
        <header className="topbar">
          <button
            type="button"
            className="icon-button"
            aria-label={sidebarOpen ? 'Close menu' : 'Open menu'}
            onClick={() => setSidebarOpen((open) => !open)}
          >
            <span className="material-symbols-rounded">{sidebarOpen ? 'close' : 'menu'}</span>
          </button>
          <div className="topbar__title">
            <span className="topbar__product">Dremio Query Doctor</span>
            {version && <span className="topbar__version">{version}</span>}
          </div>
        </header>

        <div className="main__content">
          <SectionRenderer section={currentSection} version={version} onSubmit={handleFormSubmit} />
        </div>
      </main>

      {loading && <Loader />}
    </div>
  );
}
