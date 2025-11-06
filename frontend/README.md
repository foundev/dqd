# DQD Frontend

Modern React-based frontend for Dremio Query Doctor (DQD).

## Overview

This is a complete rewrite of the DQD frontend using React, TypeScript, and modern web technologies. It provides an intuitive interface for analyzing Dremio query profiles and system metrics.

## Tech Stack

- **React 18** - UI framework
- **TypeScript** - Type safety
- **Vite** - Fast build tool and dev server
- **React Router** - Client-side routing
- **TanStack Query** - Server state management
- **Axios** - HTTP client
- **Tailwind CSS** - Utility-first CSS framework

## Project Structure

```
frontend/
├── src/
│   ├── api/           # API client and endpoints
│   ├── components/    # Reusable UI components
│   ├── hooks/         # Custom React hooks
│   ├── pages/         # Page components
│   ├── types/         # TypeScript type definitions
│   ├── App.tsx        # Main app component
│   ├── main.tsx       # Application entry point
│   └── index.css      # Global styles
├── public/            # Static assets
├── index.html         # HTML template
├── package.json       # Dependencies and scripts
├── tsconfig.json      # TypeScript configuration
├── vite.config.ts     # Vite configuration
└── tailwind.config.js # Tailwind CSS configuration
```

## Prerequisites

- Node.js 18+ or higher
- npm, yarn, or pnpm

## Quick Start

### 1. Install Dependencies

```bash
npm install
# or
yarn install
# or
pnpm install
```

### 2. Start Development Server

```bash
npm run dev
# or
yarn dev
# or
pnpm dev
```

The app will start on `http://localhost:3000`

### 3. Build for Production

```bash
npm run build
# or
yarn build
# or
pnpm build
```

## Configuration

### Environment Variables

Create a `.env` file in the root directory:

```env
# API URL - defaults to /api (proxied to backend)
VITE_API_URL=/api

# For Go backend on different port:
# VITE_API_URL=http://localhost:8080/api
```

### Development Proxy

The Vite dev server is configured to proxy `/api` requests to `http://localhost:8080` by default. This allows you to develop the frontend against either the Java or Go backend without CORS issues.

To change the proxy target, edit `vite.config.ts`:

```typescript
server: {
  port: 3000,
  proxy: {
    '/api': {
      target: 'http://localhost:8080', // Change this
      changeOrigin: true,
    },
  },
},
```

## Available Scripts

- `npm run dev` - Start development server
- `npm run build` - Build for production
- `npm run preview` - Preview production build
- `npm run lint` - Run ESLint

## Features

### Implemented

- ✅ Modern, responsive UI with Tailwind CSS
- ✅ Client-side routing with React Router
- ✅ API client with Axios and React Query
- ✅ File upload component
- ✅ Version display from `/about.json`
- ✅ Simple profile analysis page

### Coming Soon

- 📋 Detailed profile analysis
- 📋 Profile comparison
- 📋 Queries.json analysis
- 📋 Schema generation tool
- 📋 IOStat visualization
- 📋 Thread top analysis
- 📋 Dark mode support
- 📋 Results caching
- 📋 Download reports

## Pages

### Home (`/`)
Landing page with overview of all available tools.

### Simple Profile Analysis (`/profile`)
Upload and analyze a single profile.json file.

### Detailed Profile Analysis (`/profile-detailed`)
In-depth analysis with operator metrics and visualizations.

### Profile Comparison (`/profiles-comparison`)
Side-by-side comparison of two profile.json files.

### Queries.json Analysis (`/queries-json`)
Bulk analysis of queries with time-window filtering.

### Schema Generation (`/schema`)
Generate reproduction scripts from profiles.

### IOStat Analysis (`/iostat`)
Visualize disk I/O statistics.

### Thread Top Analysis (`/top`)
Analyze Linux top output for thread performance.

## API Integration

The frontend communicates with the backend through a well-defined API client (`src/api/client.ts`). Each API module corresponds to a backend endpoint:

- `aboutApi` - Version information
- `profileApi` - Profile analysis
- `queriesApi` - Queries analysis
- `reproApi` - Schema generation
- `systemApi` - System analysis (iostat, top)

## Styling

The app uses Tailwind CSS with a custom theme matching the original DQD colors:

```javascript
colors: {
  primary: {
    DEFAULT: '#006493',
    light: '#cae6ff',
    dark: '#001e30',
  },
  // ...
}
```

## Development Tips

### Hot Module Replacement (HMR)

Vite provides instant HMR, so your changes will be reflected immediately in the browser without full page reloads.

### TypeScript

The project uses strict TypeScript settings. Always define types for your components, props, and API responses.

### React Query

Use React Query for all API calls. It provides automatic caching, refetching, and error handling:

```typescript
const { data, isLoading, error } = useQuery({
  queryKey: ['profile', id],
  queryFn: () => profileApi.analyzeProfile(file),
});
```

## Browser Support

- Chrome/Edge (latest 2 versions)
- Firefox (latest 2 versions)
- Safari (latest 2 versions)

## Contributing

1. Follow the existing code style
2. Use TypeScript for all new code
3. Add proper types for all props and API responses
4. Test in multiple browsers
5. Update documentation as needed

## License

Apache License 2.0 - See LICENSE file for details
