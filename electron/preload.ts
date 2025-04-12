const { contextBridge, ipcRenderer } = require('electron');

contextBridge.exposeInMainWorld('electronAPI', {
    getOptions: () => ipcRenderer.invoke('get-options'),
});
