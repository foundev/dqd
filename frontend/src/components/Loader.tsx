export function Loader() {
  return (
    <div className="loader-overlay" role="alert" aria-live="assertive">
      <div className="loader-card">
        <div className="spinner" aria-hidden="true" />
        <h3>Analysis in progress…</h3>
        <p>Your reports are being generated. This can take a little while.</p>
      </div>
    </div>
  );
}
