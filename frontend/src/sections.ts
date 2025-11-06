import { Section } from './types';

const formatDateForInput = (date: Date): string => {
  const tzOffset = date.getTimezoneOffset() * 60000;
  const localISOTime = new Date(date.getTime() - tzOffset).toISOString();
  return localISOTime.split('T')[0];
};

const oneHundredTwentyDaysAgo = (): string => {
  const date = new Date();
  date.setDate(date.getDate() - 120);
  return formatDateForInput(date);
};

const tomorrowDate = (): string => {
  const date = new Date();
  date.setDate(date.getDate() + 1);
  return formatDateForInput(date);
};

export const SECTIONS: Section[] = [
  {
    id: 'home',
    navLabel: 'Home',
    navIcon: 'dashboard',
    hero: {
      eyebrow: 'Welcome',
      title: (version) => version || 'Dremio Query Doctor',
      subtitle:
        'Bring every profile.json, queries.json and related diagnostic artifact into a single guided experience. Upload your files, let DQD do the heavy lifting, and accelerate your investigation.'
    },
    cards: [
      {
        icon: 'rocket_launch',
        title: 'Insights in minutes',
        body: 'Upload support bundles and receive readable summaries that speed up triage and decision making.'
      },
      {
        icon: 'travel_explore',
        title: 'Use-case navigation',
        body: 'Switch between simple and detailed profiles, comparisons, queries.json analysis, IOStat, and more.'
      },
      {
        icon: 'hub',
        title: 'Guided workflow',
        body: 'Each section explains its purpose and the files you need in order to get the best report.'
      },
      {
        icon: 'code',
        title: 'Open for contributions',
        body: 'Want to improve DQD? Share ideas or submit pull requests directly on GitHub.'
      }
    ]
  },
  {
    id: 'profile',
    navLabel: 'Simple Profile',
    navIcon: 'table',
    hero: {
      eyebrow: 'Profile JSON',
      title: () => 'Simple profile.json analysis',
      subtitle: 'Generate a filterable and exportable operator view to surface hotspots quickly.'
    },
    form: {
      id: 'simple-profile-form',
      title: 'Upload a profile',
      description: 'Provide a raw or compressed profile.json to create a concise report that is easy to share.',
      action: '/simple-profile',
      method: 'POST',
      encType: 'multipart/form-data',
      buttonText: 'Analyze profile',
      buttonIcon: 'play_arrow',
      fields: [
        {
          type: 'file',
          name: 'profile1',
          label: 'Profile file',
          helper: 'Accepted formats: .tar, .gz, .tgz, .zip, .json',
          accept: '.tar, .gz, .tgz, .zip, .json',
          requiredCount: 1,
          fullWidth: true,
          placeholder: 'No file selected',
          buttonLabel: 'Browse'
        }
      ]
    }
  },
  {
    id: 'profile-detailed',
    navLabel: 'Detailed Profile',
    navIcon: 'insights',
    hero: {
      eyebrow: 'Advanced analysis',
      title: () => 'Detailed profile.json analysis',
      subtitle: 'Dig into operator-level insights: row estimates, per-node resource usage and critical hotspots.'
    },
    form: {
      id: 'detailed-profile-form',
      title: 'Run a detailed analysis',
      description: 'Upload the full profile.json to pinpoint costly operators and resource imbalances.',
      action: '/profile',
      method: 'POST',
      encType: 'multipart/form-data',
      buttonText: 'Analyze in depth',
      buttonIcon: 'analytics',
      fields: [
        {
          type: 'file',
          name: 'profile1',
          label: 'Profile file',
          helper: 'Accepted formats: .tar, .gz, .tgz, .zip, .json',
          accept: '.tar, .gz, .tgz, .zip, .json',
          requiredCount: 1,
          fullWidth: true,
          placeholder: 'No file selected',
          buttonLabel: 'Browse'
        }
      ]
    }
  },
  {
    id: 'profiles-comparison',
    navLabel: 'Profile Comparison',
    navIcon: 'compare_arrows',
    hero: {
      eyebrow: 'Diff and evolution',
      title: () => 'Compare two profiles',
      subtitle: 'Measure execution plan changes, highlight regressions, and spot operator differences between runs.'
    },
    form: {
      id: 'detailed-profiles-form',
      title: 'Compare two executions',
      description: 'Upload two profile files to generate a difference report that emphasises meaningful changes.',
      action: '/profiles',
      method: 'POST',
      encType: 'multipart/form-data',
      buttonText: 'Compare profiles',
      buttonIcon: 'sync_alt',
      fields: [
        {
          type: 'file',
          name: 'compare_profile',
          label: 'Profiles to compare',
          helper: 'Select exactly two files: .tar, .gz, .tgz, .zip, .json',
          accept: '.tar, .gz, .tgz, .zip, .json',
          requiredCount: 2,
          multiple: true,
          fullWidth: true,
          placeholder: 'No file selected',
          buttonLabel: 'Select files'
        }
      ]
    }
  },
  {
    id: 'queries-json',
    navLabel: 'Queries JSON',
    navIcon: 'equalizer',
    hero: {
      eyebrow: 'Query analytics',
      title: () => 'Analyze queries.json',
      subtitle: 'Summarise days or weeks of workload: discover bottlenecks, heavy sessions and failure patterns.'
    },
    form: {
      id: 'detailed-queries-form',
      title: 'Configure the analysis',
      description: 'Optionally narrow the time range and result limits to focus on the most relevant queries.',
      action: '/queriesjson',
      method: 'POST',
      encType: 'multipart/form-data',
      buttonText: 'Analyze queries',
      buttonIcon: 'query_stats',
      fields: [
        {
          type: 'date',
          name: 'start_date',
          label: 'Start date',
          helper: 'Filter queries that start after this date',
          defaultValue: () => oneHundredTwentyDaysAgo()
        },
        {
          type: 'time',
          name: 'start_time',
          label: 'Start time',
          helper: 'Filter queries that start after this time',
          defaultValue: '00:00'
        },
        {
          type: 'date',
          name: 'end_date',
          label: 'End date',
          helper: 'Filter queries that finish before this date',
          defaultValue: () => tomorrowDate()
        },
        {
          type: 'time',
          name: 'end_time',
          label: 'End time',
          helper: 'Filter queries that finish before this time',
          defaultValue: '00:00'
        },
        {
          type: 'select',
          name: 'limit',
          label: 'Maximum per category',
          helper: 'Larger limits increase report detail but may take longer to generate.',
          defaultValue: '5',
          options: [
            { value: '1', label: '1' },
            { value: '5', label: '5' },
            { value: '25', label: '25' },
            { value: '100', label: '100' },
            { value: '1000', label: '1000' }
          ]
        },
        {
          type: 'select',
          name: 'window',
          label: 'Aggregation window',
          helper: 'Choose the time bucket size used for temporal aggregations.',
          defaultValue: '60000',
          options: [
            { value: '1000', label: '1 second' },
            { value: '60000', label: '1 minute' },
            { value: '86400000', label: '1 day' }
          ]
        },
        {
          type: 'hidden',
          name: 'query_report_type',
          defaultValue: 'INTERACTIVE'
        },
        {
          type: 'file',
          name: 'queriesjson',
          label: 'queries.json archive',
          helper: 'Accepted formats: .tar, .tar.gz, .tgz, .tar.xz, .tar.bzip2, .bzip2, .gz, .zip, .json',
          accept: '.tar, .tar.gz, .tgz, .tar.xz, .tar.bzip2, .bzip2, .gz, .zip, .json',
          requiredCount: 1,
          fullWidth: true,
          placeholder: 'No file selected',
          buttonLabel: 'Select file'
        }
      ]
    }
  },
  {
    id: 'schema',
    navLabel: 'Schema & Reproduction',
    navIcon: 'schema',
    hero: {
      eyebrow: 'Reproduction',
      title: () => 'Schema generator and scripts',
      subtitle: 'Produce ready-to-run scripts to recreate sources, PDS and VDS referenced in a profile.json.'
    },
    form: {
      id: 'schema-form',
      title: 'Prepare schema generation',
      description: 'Customize dataset sizes and creation parameters before submitting your profile.json.',
      action: '/reproduction',
      method: 'POST',
      encType: 'multipart/form-data',
      buttonText: 'Generate scripts',
      buttonIcon: 'construction',
      fields: [
        {
          type: 'text',
          inputType: 'number',
          name: 'records',
          label: 'Rows per PDS',
          helper: 'Used when creating sample tables.',
          defaultValue: '20'
        },
        {
          type: 'text',
          inputType: 'number',
          name: 'timeout',
          label: 'Timeout (seconds)',
          helper: 'Maximum time allowed while creating objects.',
          defaultValue: '60'
        },
        {
          type: 'select',
          name: 'defaultCtasFormat',
          label: 'Default CTAS format',
          helper: 'Optional: override the default CTAS format when generating tables.',
          defaultValue: '',
          options: [
            { value: '', label: 'Default' },
            { value: 'ICEBERG', label: 'ICEBERG' },
            { value: 'PARQUET', label: 'PARQUET' }
          ]
        },
        {
          type: 'text',
          name: 'nasPath',
          label: 'Base source directory',
          helper: 'Optional: path accessible from all nodes.'
        },
        {
          type: 'textarea',
          name: 'columnDefYaml',
          label: 'Column override YAML',
          helper: 'Optional: specify exact values per column. See the example below.',
          placeholder: '# Add your custom column definitions here',
          fullWidth: true,
          rows: 8
        },
        {
          type: 'file',
          name: 'profile',
          label: 'Profile.json source',
          helper: 'Accepted formats: .tar, .gz, .tgz, .zip, .json',
          accept: '.tar, .gz, .tgz, .zip, .json',
          requiredCount: 1,
          fullWidth: true,
          placeholder: 'No file selected',
          buttonLabel: 'Select file'
        }
      ]
    },
    checklist: [
      'Upload your profile.json (zip or raw) and wait for the generated archive.',
      'Extract the archive and switch to the generated directory.',
      'Run `bash create.sh --host "http://localhost:9047" -u "myuser" -p "mypass"` to recreate the objects.',
      'Review `debug.sql` if a step fails to identify problematic statements.',
      'Create any missing sources manually if required, then rerun the script.'
    ],
    codeSample: {
      title: 'Sample column override YAML',
      code: `tables:\n  - name: '"ns 1".table1'\n    columns:\n      - name: status\n        values: [Active, Inactive, Suspended]\n      - name: customerType\n        values: [Silver, Gold, Platinum]`
    }
  },
  {
    id: 'iostat-analysis',
    navLabel: 'IOStat Analysis',
    navIcon: 'monitoring',
    hero: {
      eyebrow: 'System',
      title: () => 'Analyze IOStat output',
      subtitle: 'Detect disk saturation and system anomalies from `iostat -x -c -d -t 1 600` output.'
    },
    form: {
      id: 'iostat-form',
      title: 'Upload IOStat output',
      description: 'The report highlights wait times, latency and per-device saturation.',
      action: '/iostat',
      method: 'POST',
      encType: 'multipart/form-data',
      buttonText: 'Analyze IOStat',
      buttonIcon: 'monitor_heart',
      fields: [
        {
          type: 'file',
          name: 'iostatfile',
          label: 'IOStat file',
          helper: 'Ideally captured with: iostat -x -c -d -t 1 600',
          accept: '.txt, .log, .out',
          requiredCount: 1,
          fullWidth: true,
          placeholder: 'No file selected',
          buttonLabel: 'Select file'
        }
      ]
    }
  },
  {
    id: 'top-analysis',
    navLabel: 'Threaded Top',
    navIcon: 'vertical_align_top',
    hero: {
      eyebrow: 'CPU and threads',
      title: () => 'Analyze threaded top output',
      subtitle: 'Use the DDC command `LINES=100 top -H -n 120 -p 1 -d 2 -bw` to capture dominant threads.'
    },
    form: {
      id: 'top-form',
      title: 'Upload threaded top file',
      description: 'Identify the busiest threads and understand how workload evolves over time.',
      action: '/ttop',
      method: 'POST',
      encType: 'multipart/form-data',
      buttonText: 'Analyze file',
      buttonIcon: 'stacked_bar_chart',
      fields: [
        {
          type: 'file',
          name: 'ttop',
          label: 'top capture',
          helper: 'Recommended flags: LINES=100 top -H -n 120 -p 1 -d 2 -bw',
          accept: '.txt, .log, .out',
          requiredCount: 1,
          fullWidth: true,
          placeholder: 'No file selected',
          buttonLabel: 'Select file'
        }
      ]
    }
  }
];
