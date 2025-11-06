import type { Section } from '../types';
import { UploadForm } from './UploadForm';

type SectionRendererProps = {
  section?: Section;
  version: string;
  onSubmit?: () => void;
};

const resolveContent = (
  value: string | ((version: string) => string) | undefined,
  version: string
): string | undefined => {
  if (!value) {
    return undefined;
  }
  if (typeof value === 'function') {
    return value(version);
  }
  return value;
};

export function SectionRenderer({ section, version, onSubmit }: SectionRendererProps) {
  if (!section) {
    return null;
  }

  const hero = section.hero;
  const title = resolveContent(hero?.title, version);
  const subtitle = resolveContent(hero?.subtitle, version);

  return (
    <section className="section" id={section.id} aria-labelledby={`${section.id}-title`}>
      <div className="section__hero">
        {hero?.eyebrow && <p className="section__eyebrow">{hero.eyebrow}</p>}
        {title && <h1 id={`${section.id}-title`}>{title}</h1>}
        {subtitle && <p>{subtitle}</p>}
      </div>

      {section.cards?.length ? (
        <div className="feature-grid">
          {section.cards.map((card) => (
            <article className="feature-card" key={card.title}>
              {card.icon && (
                <span className="feature-card__icon material-symbols-rounded">{card.icon}</span>
              )}
              <h3>{card.title}</h3>
              <p>{card.body}</p>
            </article>
          ))}
        </div>
      ) : null}

      {section.form && <UploadForm formConfig={section.form} onSubmit={onSubmit} />}

      {section.checklist?.length ? (
        <div className="section__checklist">
          <h2>Suivez les étapes</h2>
          <ol>
            {section.checklist.map((item, index) => (
              <li key={index}>{item}</li>
            ))}
          </ol>
        </div>
      ) : null}

      {section.codeSample ? (
        <div className="code-card">
          <div className="code-card__header">
            <span className="material-symbols-rounded">description</span>
            <span>{section.codeSample.title}</span>
          </div>
          <pre>
            <code>{section.codeSample.code}</code>
          </pre>
        </div>
      ) : null}
    </section>
  );
}
