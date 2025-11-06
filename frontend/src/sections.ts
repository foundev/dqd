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
    navLabel: 'Accueil',
    navIcon: 'dashboard',
    hero: {
      eyebrow: 'Bienvenue',
      title: (version) => version || 'Dremio Query Doctor',
      subtitle:
        'Retrouvez en un seul lieu des diagnostics rapides pour vos fichiers profile.json, queries.json et autres artefacts de support. Téléversez vos fichiers, laissez DQD travailler et accélérez vos analyses.'
    },
    cards: [
      {
        icon: 'rocket_launch',
        title: 'Analyses en quelques minutes',
        body: 'Déposez vos fichiers collectés et obtenez des rapports lisibles pour accélérer les investigations.'
      },
      {
        icon: 'travel_explore',
        title: 'Vue d’ensemble claire',
        body: 'Naviguez par thématique : profils simples ou détaillés, comparaisons, requêtes, IOstat et plus encore.'
      },
      {
        icon: 'hub',
        title: 'Un workflow guidé',
        body: 'Chaque page rappelle le but de l’outil et ce qu’il faut fournir pour tirer le meilleur du diagnostic.'
      },
      {
        icon: 'code',
        title: 'Toujours extensible',
        body: 'Envie de contribuer ? Le projet est ouvert : remontez vos idées et améliorations directement sur GitHub.'
      }
    ]
  },
  {
    id: 'profile',
    navLabel: 'Profil simple',
    navIcon: 'table',
    hero: {
      eyebrow: 'Profil JSON',
      title: () => 'Analyse simple de profile.json',
      subtitle: 'Obtenez une vue exportable et filtrable des opérateurs afin d’identifier rapidement les points chauds.'
    },
    form: {
      id: 'simple-profile-form',
      title: 'Téléverser un profil',
      description:
        'Ajoutez un fichier profile.json (brut ou compressé) pour générer un rapport synthétique, prêt à être partagé.',
      action: '/simple-profile',
      method: 'POST',
      encType: 'multipart/form-data',
      buttonText: 'Analyser le profil',
      buttonIcon: 'play_arrow',
      fields: [
        {
          type: 'file',
          name: 'profile1',
          label: 'Fichier à analyser',
          helper: 'Formats acceptés : .tar, .gz, .tgz, .zip, .json',
          accept: '.tar, .gz, .tgz, .zip, .json',
          requiredCount: 1,
          fullWidth: true,
          placeholder: 'Aucun fichier sélectionné',
          buttonLabel: 'Choisir un fichier'
        }
      ]
    }
  },
  {
    id: 'profile-detailed',
    navLabel: 'Profil détaillé',
    navIcon: 'insights',
    hero: {
      eyebrow: 'Profil avancé',
      title: () => 'Analyse détaillée de profile.json',
      subtitle:
        'Approfondissez l’étude des opérateurs : estimations de lignes, consommation par nœud et ressources critiques.'
    },
    form: {
      id: 'detailed-profile-form',
      title: 'Lancer une analyse détaillée',
      description:
        'Identifiez les opérateurs les plus coûteux et les éventuels déséquilibres de ressources en fournissant un profile.json complet.',
      action: '/profile',
      method: 'POST',
      encType: 'multipart/form-data',
      buttonText: 'Analyser en profondeur',
      buttonIcon: 'analytics',
      fields: [
        {
          type: 'file',
          name: 'profile1',
          label: 'Fichier profil',
          helper: 'Formats acceptés : .tar, .gz, .tgz, .zip, .json',
          accept: '.tar, .gz, .tgz, .zip, .json',
          requiredCount: 1,
          fullWidth: true,
          placeholder: 'Aucun fichier sélectionné',
          buttonLabel: 'Joindre un profile.json'
        }
      ]
    }
  },
  {
    id: 'profiles-comparison',
    navLabel: 'Comparer des profils',
    navIcon: 'compare_arrows',
    hero: {
      eyebrow: 'Diff & évolution',
      title: () => 'Comparaison de deux profiles',
      subtitle:
        'Mesurez l’évolution des plans d’exécution, repérez les régressions de performance et les différences d’opérateur.'
    },
    form: {
      id: 'detailed-profiles-form',
      title: 'Comparer deux exécutions',
      description:
        'Téléversez deux profils pour générer un rapport de différences mettant en évidence les changements majeurs.',
      action: '/profiles',
      method: 'POST',
      encType: 'multipart/form-data',
      buttonText: 'Comparer',
      buttonIcon: 'sync_alt',
      fields: [
        {
          type: 'file',
          name: 'compare_profile',
          label: 'Profils à comparer',
          helper: 'Sélectionnez exactement deux fichiers : .tar, .gz, .tgz, .zip, .json',
          accept: '.tar, .gz, .tgz, .zip, .json',
          requiredCount: 2,
          multiple: true,
          fullWidth: true,
          placeholder: 'Aucun fichier sélectionné',
          buttonLabel: 'Choisir deux fichiers'
        }
      ]
    }
  },
  {
    id: 'queries-json',
    navLabel: 'Queries.json',
    navIcon: 'equalizer',
    hero: {
      eyebrow: 'Analyse requêtes',
      title: () => 'Étude de queries.json',
      subtitle:
        'Synthétisez des dizaines de jours d’activité : débordements, goulots d’étranglement, sessions les plus coûteuses.'
    },
    form: {
      id: 'detailed-queries-form',
      title: 'Configurer l’analyse de requêtes',
      description:
        'Limitez éventuellement la plage temporelle et le nombre de résultats pour générer un rapport ciblé.',
      action: '/queriesjson',
      method: 'POST',
      encType: 'multipart/form-data',
      buttonText: 'Analyser les requêtes',
      buttonIcon: 'query_stats',
      fields: [
        {
          type: 'date',
          name: 'start_date',
          label: 'Date de début',
          helper: 'Filtrer les requêtes qui commencent après cette date',
          defaultValue: () => oneHundredTwentyDaysAgo()
        },
        {
          type: 'time',
          name: 'start_time',
          label: 'Heure de début',
          helper: 'Filtrer les requêtes qui commencent après cette heure',
          defaultValue: '00:00'
        },
        {
          type: 'date',
          name: 'end_date',
          label: 'Date de fin',
          helper: 'Filtrer les requêtes qui se terminent avant cette date',
          defaultValue: () => tomorrowDate()
        },
        {
          type: 'time',
          name: 'end_time',
          label: 'Heure de fin',
          helper: 'Filtrer les requêtes qui finissent avant cette heure',
          defaultValue: '00:00'
        },
        {
          type: 'select',
          name: 'limit',
          label: 'Nombre maximum par catégorie',
          helper:
            'Plus la limite est élevée, plus le rapport est détaillé (et potentiellement plus long à générer).',
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
          label: 'Taille de fenêtre',
          helper: 'Définissez la granularité des agrégats temporels.',
          defaultValue: '60000',
          options: [
            { value: '1000', label: '1 seconde' },
            { value: '60000', label: '1 minute' },
            { value: '86400000', label: '1 jour' }
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
          label: 'Archive queries.json',
          helper:
            'Formats acceptés : .tar, .tar.gz, .tgz, .tar.xz, .tar.bzip2, .bzip2, .gz, .zip, .json',
          accept: '.tar, .tar.gz, .tgz, .tar.xz, .tar.bzip2, .bzip2, .gz, .zip, .json',
          requiredCount: 1,
          fullWidth: true,
          placeholder: 'Aucun fichier sélectionné',
          buttonLabel: 'Choisir un fichier queries.json'
        }
      ]
    }
  },
  {
    id: 'schema',
    navLabel: 'Schéma & repro',
    navIcon: 'schema',
    hero: {
      eyebrow: 'Reproduction',
      title: () => 'Générateur de schémas & scripts',
      subtitle:
        'Produisez des scripts d’initialisation pour recréer les sources, PDS et VDS observés dans un profile.json.'
    },
    form: {
      id: 'schema-form',
      title: 'Préparer la génération de schémas',
      description:
        'Personnalisez la taille des jeux de données et les paramètres de création avant de déposer votre profile.json.',
      action: '/reproduction',
      method: 'POST',
      encType: 'multipart/form-data',
      buttonText: 'Générer les scripts',
      buttonIcon: 'construction',
      fields: [
        {
          type: 'text',
          inputType: 'number',
          name: 'records',
          label: 'Nombre de lignes par PDS',
          helper: 'Utilisé pour créer des échantillons de tables.',
          defaultValue: '20'
        },
        {
          type: 'text',
          inputType: 'number',
          name: 'timeout',
          label: 'Timeout (secondes)',
          helper: 'Temps maximal alloué à la création des objets.',
          defaultValue: '60'
        },
        {
          type: 'select',
          name: 'defaultCtasFormat',
          label: 'Format CTAS par défaut',
          helper: 'Optionnel : force un format spécifique lors des créations CTAS.',
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
          label: 'Répertoire source',
          helper: 'Optionnel : chemin accessible depuis l’ensemble des nœuds.'
        },
        {
          type: 'textarea',
          name: 'columnDefYaml',
          label: 'Overrides de colonnes (YAML)',
          helper:
            'Facultatif : définir des valeurs précises par colonne. Voir exemple ci-dessous.',
          placeholder: '# Ajoutez ici vos définitions personnalisées',
          fullWidth: true,
          rows: 8
        },
        {
          type: 'file',
          name: 'profile',
          label: 'Profile.json source',
          helper: 'Formats acceptés : .tar, .gz, .tgz, .zip, .json',
          accept: '.tar, .gz, .tgz, .zip, .json',
          requiredCount: 1,
          fullWidth: true,
          placeholder: 'Aucun fichier sélectionné',
          buttonLabel: 'Joindre un profile.json'
        }
      ]
    },
    checklist: [
      'Téléversez le profile.json (zip ou fichier brut) et patientez jusqu’à réception de l’archive générée.',
      'Décompressez l’archive et positionnez-vous dans le dossier créé.',
      'Lancez le script `bash create.sh --host "http://localhost:9047" -u "monuser" -p "monmdp"` pour recréer les objets.',
      'Consultez `debug.sql` si une étape échoue afin d’identifier les requêtes problématiques.',
      'Créez manuellement les sources manquantes si nécessaire, puis rejouez le script.'
    ],
    codeSample: {
      title: 'Exemple de surcharge YAML',
      code: `tables:\n  - name: '"ns 1".table1'\n    columns:\n      - name: status\n        values: [Active, Inactive, Suspended]\n      - name: customerType\n        values: [Silver, Gold, Platinum]`
    }
  },
  {
    id: 'iostat-analysis',
    navLabel: 'Analyse IOStat',
    navIcon: 'monitoring',
    hero: {
      eyebrow: 'Système',
      title: () => 'Analyser la sortie IOStat',
      subtitle:
        'Repérez les saturations de disques et anomalies système en important la sortie `iostat -x -c -d -t 1 600`.'
    },
    form: {
      id: 'iostat-form',
      title: 'Téléverser votre sortie IOStat',
      description:
        'Le rapport mettra en avant les pics d’attente disque, la latence et la saturation par device.',
      action: '/iostat',
      method: 'POST',
      encType: 'multipart/form-data',
      buttonText: 'Analyser IOStat',
      buttonIcon: 'monitor_heart',
      fields: [
        {
          type: 'file',
          name: 'iostatfile',
          label: 'Fichier IOStat',
          helper: 'Utilisez idéalement la commande : iostat -x -c -d -t 1 600',
          accept: '.txt, .log, .out',
          requiredCount: 1,
          fullWidth: true,
          placeholder: 'Aucun fichier sélectionné',
          buttonLabel: 'Sélectionner un fichier'
        }
      ]
    }
  },
  {
    id: 'top-analysis',
    navLabel: 'Threaded Top',
    navIcon: 'vertical_align_top',
    hero: {
      eyebrow: 'CPU & Threads',
      title: () => 'Analyse threaded top',
      subtitle:
        'Appelez ce module avec la sortie DDC `LINES=100 top -H -n 120 -p 1 -d 2 -bw` pour déceler les threads dominants.'
    },
    form: {
      id: 'top-form',
      title: 'Fichier threaded top',
      description:
        'Identifiez les threads les plus consommateurs et les évolutions de charge de votre service.',
      action: '/ttop',
      method: 'POST',
      encType: 'multipart/form-data',
      buttonText: 'Analyser le fichier',
      buttonIcon: 'stacked_bar_chart',
      fields: [
        {
          type: 'file',
          name: 'ttop',
          label: 'Capture top',
          helper: 'Recommandé : LINES=100 top -H -n 120 -p 1 -d 2 -bw',
          accept: '.txt, .log, .out',
          requiredCount: 1,
          fullWidth: true,
          placeholder: 'Aucun fichier sélectionné',
          buttonLabel: 'Sélectionner un fichier'
        }
      ]
    }
  }
];
