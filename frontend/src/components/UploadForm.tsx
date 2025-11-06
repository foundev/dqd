import { useCallback, useMemo, useState } from 'react';
import type { FileField, FormConfig, FormField, ValueFactory } from '../types';

const resolveValue = (value?: ValueFactory): string | undefined => {
  if (typeof value === 'function') {
    return value();
  }
  return value;
};

const isFileField = (field: FormField): field is FileField => field.type === 'file';

type UploadFormProps = {
  formConfig: FormConfig;
  onSubmit?: () => void;
};

export function UploadForm({ formConfig, onSubmit }: UploadFormProps) {
  const fileFields = useMemo(() => formConfig.fields.filter(isFileField), [formConfig.fields]);

  const [fileState, setFileState] = useState<Record<string, boolean>>(() => {
    const initialState: Record<string, boolean> = {};
    fileFields.forEach((field) => {
      initialState[field.name] = false;
    });
    return initialState;
  });

  const [fileSummaries, setFileSummaries] = useState<Record<string, string>>(() => {
    const initialSummaries: Record<string, string> = {};
    fileFields.forEach((field) => {
      initialSummaries[field.name] = field.placeholder || 'Aucun fichier sélectionné';
    });
    return initialSummaries;
  });

  const createFileChangeHandler = useCallback(
    (field: FileField) => (event: React.ChangeEvent<HTMLInputElement>) => {
      const files = Array.from(event.target.files ?? []);
      const hasEnoughFiles = field.requiredCount ? files.length >= field.requiredCount : files.length > 0;

      setFileState((prev) => ({
        ...prev,
        [field.name]: hasEnoughFiles
      }));

      setFileSummaries((prev) => ({
        ...prev,
        [field.name]: files.length
          ? files.map((file) => file.name).join(', ')
          : field.placeholder || 'Aucun fichier sélectionné'
      }));
    },
    []
  );

  const allFilesReady = fileFields.length === 0 || Object.values(fileState).every(Boolean);

  const hiddenFields = formConfig.fields.filter((field) => field.type === 'hidden');
  const visibleFields = formConfig.fields.filter((field) => field.type !== 'hidden');

  return (
    <form
      id={formConfig.id}
      className="upload-form"
      action={formConfig.action}
      method={formConfig.method}
      encType={formConfig.encType}
      onSubmit={() => onSubmit?.()}
    >
      <div className="upload-form__header">
        {formConfig.title && <h2>{formConfig.title}</h2>}
        {formConfig.description && <p>{formConfig.description}</p>}
      </div>

      <div className="form-grid">
        {visibleFields.map((field) => {
          const fieldId = `${formConfig.id}-${field.name}`;
          const className = ['form-field', field.fullWidth ? 'form-field--full' : '']
            .filter(Boolean)
            .join(' ');

          if (field.type === 'file') {
            return (
              <div className={className} key={field.name}>
                <p className="form-label">
                  <span className="material-symbols-rounded">attach_file</span>
                  {field.label}
                </p>
                <label className="file-upload">
                  <span className="file-upload__icon material-symbols-rounded">file_upload</span>
                  <div className="file-upload__text">
                    <span className="file-upload__title">{field.buttonLabel || 'Choisir un fichier'}</span>
                    <span className="file-upload__subtitle">{fileSummaries[field.name]}</span>
                  </div>
                  <input
                    id={fieldId}
                    className="file-upload__input"
                    type="file"
                    name={field.name}
                    accept={field.accept}
                    multiple={field.multiple ?? (field.requiredCount ?? 0) > 1}
                    onChange={createFileChangeHandler(field)}
                    required={field.required !== false}
                  />
                </label>
                {field.helper && <small className="form-helper">{field.helper}</small>}
              </div>
            );
          }

          if (field.type === 'select') {
            return (
              <div className={className} key={field.name}>
                <label className="form-label" htmlFor={fieldId}>
                  <span className="material-symbols-rounded">tune</span>
                  {field.label}
                </label>
                <select
                  id={fieldId}
                  name={field.name}
                  defaultValue={resolveValue(field.defaultValue)}
                  required={field.required !== false}
                >
                  {field.options?.map((option) => (
                    <option key={option.value} value={option.value}>
                      {option.label}
                    </option>
                  ))}
                </select>
                {field.helper && <small className="form-helper">{field.helper}</small>}
              </div>
            );
          }

          if (field.type === 'textarea') {
            return (
              <div className={className} key={field.name}>
                <label className="form-label" htmlFor={fieldId}>
                  <span className="material-symbols-rounded">notes</span>
                  {field.label}
                </label>
                <textarea
                  id={fieldId}
                  name={field.name}
                  rows={field.rows ?? 6}
                  placeholder={field.placeholder}
                  defaultValue={resolveValue(field.defaultValue)}
                  required={field.required !== false}
                />
                {field.helper && <small className="form-helper">{field.helper}</small>}
              </div>
            );
          }

          const inputType = field.type === 'date' || field.type === 'time' ? field.type : field.inputType || 'text';

          return (
            <div className={className} key={field.name}>
              <label className="form-label" htmlFor={fieldId}>
                <span className="material-symbols-rounded">edit_square</span>
                {field.label}
              </label>
              <input
                id={fieldId}
                type={inputType}
                name={field.name}
                defaultValue={resolveValue(field.defaultValue)}
                placeholder={field.placeholder}
                required={field.required !== false}
              />
              {field.helper && <small className="form-helper">{field.helper}</small>}
            </div>
          );
        })}
      </div>

      {hiddenFields.map((field) => (
        <input
          key={field.name}
          type="hidden"
          name={field.name}
          value={resolveValue(field.defaultValue)}
        />
      ))}

      <div className="upload-form__footer">
        <button className="primary-button" type="submit" disabled={!allFilesReady}>
          {formConfig.buttonIcon && (
            <span className="material-symbols-rounded">{formConfig.buttonIcon}</span>
          )}
          <span>{formConfig.buttonText || 'Envoyer'}</span>
        </button>
      </div>
    </form>
  );
}
