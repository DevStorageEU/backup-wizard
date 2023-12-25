import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TopHeaderbarComponent } from './top-headerbar.component';

describe('TopHeaderbarComponent', () => {
  let component: TopHeaderbarComponent;
  let fixture: ComponentFixture<TopHeaderbarComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ TopHeaderbarComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(TopHeaderbarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
