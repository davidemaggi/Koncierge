import { app, BrowserWindow, ipcMain } from 'electron';
import * as path from 'path';
import { fetchOptions } from '../backend/apiService';

let win: BrowserWindow;

async function createWindow() {
    win = new BrowserWindow({
        width: 800,
        height: 600,
        webPreferences: {
            contextIsolation: true,
            preload: path.join(__dirname, 'preload.js'),
        },
    });

    await win.loadURL('http://localhost:4200'); // Angular dev server
}

app.whenReady().then(createWindow);

ipcMain.handle('get-options', async () => {
    return await fetchOptions();
});
