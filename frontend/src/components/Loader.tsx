export function Loader() {
  return (
    <div className="loader-overlay" role="alert" aria-live="assertive">
      <div className="loader-card">
        <div className="spinner" aria-hidden="true" />
        <h3>Analyse en cours…</h3>
        <p>Vos rapports se préparent. Cette opération peut prendre quelques instants.</p>
      </div>
    </div>
  );
}
