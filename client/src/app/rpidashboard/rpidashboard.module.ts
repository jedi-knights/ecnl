import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RankingComponent } from './ranking/ranking.component';
import { DivisionSelectorComponent } from './division-selector/division-selector.component';



@NgModule({
  declarations: [
    RankingComponent,
    DivisionSelectorComponent
  ],
  imports: [
    CommonModule
  ]
})
export class RPIDashboardModule { }
