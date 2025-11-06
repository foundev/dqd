export type ValueFactory = string | (() => string);

export type FieldType = 'text' | 'textarea' | 'select' | 'file' | 'date' | 'time' | 'hidden';

export interface BaseField {
  type: FieldType;
  name: string;
  label?: string;
  helper?: string;
  defaultValue?: ValueFactory;
  placeholder?: string;
  fullWidth?: boolean;
  required?: boolean;
}

export interface TextField extends BaseField {
  type: 'text' | 'date' | 'time';
  inputType?: string;
}

export interface SelectField extends BaseField {
  type: 'select';
  options: Array<{ value: string; label: string }>;
}

export interface TextareaField extends BaseField {
  type: 'textarea';
  rows?: number;
}

export interface FileField extends BaseField {
  type: 'file';
  accept?: string;
  requiredCount?: number;
  multiple?: boolean;
  buttonLabel?: string;
}

export interface HiddenField extends BaseField {
  type: 'hidden';
}

export type FormField = TextField | SelectField | TextareaField | FileField | HiddenField;

export interface FormConfig {
  id: string;
  title?: string;
  description?: string;
  action: string;
  method: string;
  encType?: string;
  buttonText?: string;
  buttonIcon?: string;
  fields: FormField[];
}

export interface HeroContent {
  eyebrow?: string;
  title?: string | ((version: string) => string);
  subtitle?: string | ((version: string) => string);
}

export interface FeatureCard {
  icon?: string;
  title: string;
  body: string;
}

export interface SectionCodeSample {
  title: string;
  code: string;
}

export interface Section {
  id: string;
  navLabel: string;
  navIcon: string;
  hero?: HeroContent;
  cards?: FeatureCard[];
  form?: FormConfig;
  checklist?: string[];
  codeSample?: SectionCodeSample;
}
