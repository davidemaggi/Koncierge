import {Component, OnInit} from '@angular/core';
import { RouterOutlet } from '@angular/router';
import {Select, SelectModule} from 'primeng/select';
import {FormsModule} from '@angular/forms';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, Select, FormsModule],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css'
})
export class AppComponent implements OnInit {
  title = 'koncierge';

  options: string[] = [];
  selectedOption: string | undefined;

  async ngOnInit() {
    this.options = await window.electronAPI.getOptions();
  }



}
