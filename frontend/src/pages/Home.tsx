export const Home = () => {
  return (
    <div className="max-w-4xl">
      <h1 className="text-3xl font-bold text-gray-900 mb-6">
        Welcome to Dremio Query Doctor
      </h1>

      <div className="bg-white rounded-lg shadow-md p-6 mb-6">
        <h2 className="text-xl font-semibold mb-4">
          Profile.json and Queries.json Analysis Tools
        </h2>
        <p className="text-gray-700 mb-4">
          DQD is a diagnostic tool for analyzing Dremio query profiles, system metrics,
          and query execution logs. Use the navigation menu to access different analysis tools.
        </p>
        <p className="text-gray-600">
          If you find something you don't like or would like to improve DQD, you can{' '}
          <a
            href="https://github.com/rsvihladremio/dqd"
            target="_blank"
            rel="noopener noreferrer"
            className="text-primary hover:text-primary-dark underline font-medium"
          >
            submit code or issues on GitHub
          </a>
          .
        </p>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div className="bg-white rounded-lg shadow-md p-6">
          <h3 className="text-lg font-semibold mb-3">📊 Profile Analysis</h3>
          <ul className="space-y-2 text-sm text-gray-700">
            <li>• Single profile.json analysis</li>
            <li>• Compare two profiles</li>
            <li>• Performance bottleneck detection</li>
            <li>• Detailed operator metrics</li>
          </ul>
        </div>

        <div className="bg-white rounded-lg shadow-md p-6">
          <h3 className="text-lg font-semibold mb-3">📈 Queries Analysis</h3>
          <ul className="space-y-2 text-sm text-gray-700">
            <li>• Bulk queries.json analysis</li>
            <li>• Time-window filtering</li>
            <li>• Query pattern detection</li>
            <li>• Performance trending</li>
          </ul>
        </div>

        <div className="bg-white rounded-lg shadow-md p-6">
          <h3 className="text-lg font-semibold mb-3">🗄️ Schema Generation</h3>
          <ul className="space-y-2 text-sm text-gray-700">
            <li>• Generate reproduction scripts</li>
            <li>• Create test datasets</li>
            <li>• SQL DDL generation</li>
            <li>• Custom column definitions</li>
          </ul>
        </div>

        <div className="bg-white rounded-lg shadow-md p-6">
          <h3 className="text-lg font-semibold mb-3">💻 System Analysis</h3>
          <ul className="space-y-2 text-sm text-gray-700">
            <li>• IOStat disk I/O analysis</li>
            <li>• Thread top analysis</li>
            <li>• Resource utilization</li>
            <li>• Performance visualizations</li>
          </ul>
        </div>
      </div>
    </div>
  );
};
