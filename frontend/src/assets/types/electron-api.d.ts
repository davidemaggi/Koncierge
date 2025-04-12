export interface ElectronAPI {
  getOptions(): Promise<string[]>;
}

declare global {
  interface Window {
    electronAPI: ElectronAPI;
  }
}
