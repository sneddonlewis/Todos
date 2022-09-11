import { Component, Inject, OnInit } from '@angular/core';

@Component({
  selector: 'app-todo-item',
  templateUrl: './todo-item.component.html',
  styleUrls: ['./todo-item.component.scss']
})
export class TodoItemComponent implements OnInit {

  public title = "item";
  public description = "a thing about the thing";

  constructor(
  ) { }

  ngOnInit(): void {
  }

}
