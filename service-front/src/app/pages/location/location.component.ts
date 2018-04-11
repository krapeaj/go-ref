import { Component, OnInit } from "@angular/core"
import { GeoLocationService } from "../../shared/geo-location.service"


@Component({
  selector: 'location',
  providers: [],
  styleUrls: [ './location.component.css' ],
  templateUrl: './location.component.html'
})
export class LocationComponent implements OnInit {

  title: string = 'My first AGM project'
  lat: number = 51.678418
  lng: number = 7.809007

  constructor(geoLocationService: GeoLocationService) {
  }

  ngOnInit(): void {
  }


}