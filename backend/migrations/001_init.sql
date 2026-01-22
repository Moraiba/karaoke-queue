CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE karaoke_sessions (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  nombre TEXT NOT NULL,
  hora_inicio TIMESTAMP NOT NULL,
  hora_fin TIMESTAMP NOT NULL,
  duracion_promedio_cancion INT NOT NULL,
  estado TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE participants (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  session_id UUID REFERENCES karaoke_sessions(id) ON DELETE CASCADE,
  nombre TEXT NOT NULL,
  cancion TEXT NOT NULL,
  video_id TEXT,
  estado TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT NOW()
);
