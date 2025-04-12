export interface ElectronAPI {
  getOptions(): Promise<string[]>;
  getPods(): Promise<string[]>;
}

declare global {
  interface Window {
    electronAPI: ElectronAPI;
  }
}
