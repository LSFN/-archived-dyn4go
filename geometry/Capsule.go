package geometry

const (
	EDGE_FEATURE_SELECTION_CRITERIA = 0.98
	EDGE_FEATURE_EXPANION_FACTOR = 0.1
)

type Capsule struct {
	radius float64
	center *Vector2
	length, capRadius float64
	foci []*Vector2
	localXAxis *Vector2
}

func (c *Capsule) GetID() string {

}

func (c *Capsule) GetCenter() *Vector2 {

}

func (c *Capsule) GetUserData() interface{} {

}

func (c *Capsule) SetUserData(data interface{}) {

}

func (c *Capsule) RotateAboutOrigin(theta float64) {

}

func (c *Capsule) RotateAboutCenter(theta float64) {

}

func (c *Capsule) RotateAboutVector2(theta float64, v *Vector2) {

}

func (c *Capsule) RotateAboutXY(theta, x, y float64) {

}

func (c *Capsule) TranslateXY(x, y float64) {

}

func (c *Capsule) TranslateVector2(v *Vector2) {

}

func (c *Capsule) Project(v *Vector2) *Interval {

}

func (c *Capsule) Contains(v *Vector2) bool {

}

func (c *Capsule) CreateAABB() *AABB {

}

func NewCapsule(width, height float64) *Capsule {
	major := width
	minor := height
	vertical := false
	if width < heigth {
		major, minor = minor, major
		vertical = true
	}

	c := new(Capsule)
	c.length = major
	c.capRadius = minor * 0.5
	c.radius = major * 0.5
	c.center = NewVector2FromXY(0, 0)

	f = (major - minor) * 0.5
	if vertical {
		c.foci = [2]*Vector2 {NewVector2FromXY(0, -f), NewVector2FromXY(0, f)}
		c.loaclXAxis = NewVector2FromXY(0, 1)
	} else {
		c.foci = [2]*Vector2 {NewVector2FromXY(-f, 0), NewVector2FromXY(f, 0)}
		c.loaclXAxis = NewVector2FromXY(1, 0)
	}

	return c
}

func (c *Capsule) GetAxes(foci []*Vector2, t *Transform) {
	if c.foci != nil {
		[]*Vector2 axes = make([]*Vector2, 2 + len(c.foci))
		axes[0] = t.GetTransformedR(c.locallXAxis)
		axes[1] = t.GetTransformedR(c.locallXAxis.GetRightHandOrthogonalVector())
		f1 := t.GetTransformed(c.foci[0])
		f2 := t.GetTransformed(c.foci[1])
		for i, f := c.foci {
			if f1.DistanceSquared(f) < f2.DistanceSquared(f) {
				axes[2+i] = f1.HereToVector2(f)
			} else {
				axes[2+i] = f2.HereToVector2(f)
			}
		}
		return axes
	} else {
		return []*Vector2{t.GetTransformedR(c.localXAxis), t.GetTransformedR(c.localXAxis).GetRightHandOrthogonalVector()}
	}
}

func (c *Capsule) GetFoci(t *Transform) {
	return []*Vector2{t.GetTransformed(c.foci[0]), t.GetTransformed(c.foci[1])}
}

func (c *Capsule) GetFarthestPoint(v *Vector2, t *Transform) {
	v.Normalize()
	
}

// -------------------------------------------------------------------------------------------------------------------------------------


	/* (non-Javadoc)
	 * @see org.dyn4j.geometry.Convex#getFarthestPoint(org.dyn4j.geometry.Vector2, org.dyn4j.geometry.Transform)
	 */
	@Override
	public Vector2 getFarthestPoint(Vector2 n, Transform transform) {
		// make sure the given direction is normalized
		n.normalize();
		// a capsule is just a radially expanded line segment
		Vector2 p = Segment.getFarthestPoint(this.foci[0], this.foci[1], n, transform);
		// apply the radial expansion
		return p.add(n.product(this.capRadius));
	}
	
	/* (non-Javadoc)
	 * @see org.dyn4j.geometry.Convex#getFarthestFeature(org.dyn4j.geometry.Vector2, org.dyn4j.geometry.Transform)
	 */
	@Override
	public Feature getFarthestFeature(Vector2 n, Transform transform) {
		// test whether the given direction is within a certain angle of the
		// local x axis. if so, use the edge feature rather than the point
		Vector2 localAxis = transform.getInverseTransformedR(n);
		Vector2 n1 = this.localXAxis.getLeftHandOrthogonalVector();
		
		// get the squared length of the localaxis and add the fudge factor
		// should always 1.0 * factor since localaxis is normalized
		double d = localAxis.dot(localAxis) * Capsule.EDGE_FEATURE_SELECTION_CRITERIA;
		// project the normal onto the localaxis normal
		double d1 = localAxis.dot(n1);
		
		// we only need to test one normal since we only care about its projection length
		// we can later determine which direction by the sign of the projection
		if (Math.abs(d1) < d) {
			// then its the farthest point
			Vector2 point = this.getFarthestPoint(n, transform);
			return new Vertex(point);
		} else {
			// compute the vector to add/sub from the foci
			Vector2 v = n1.multiply(this.capRadius);
			// compute an expansion amount based on the width of the shape
			Vector2 e = this.localXAxis.product(this.length * 0.5 * EDGE_FEATURE_EXPANSION_FACTOR);
			if (d1 > 0) {
				Vector2 p1 = this.foci[0].sum(v).subtract(e);
				Vector2 p2 = this.foci[1].sum(v).add(e);
				// return the full bottom side
				return Segment.getFarthestFeature(p1, p2, n, transform);
			} else {
				Vector2 p1 = this.foci[0].difference(v).subtract(e);
				Vector2 p2 = this.foci[1].difference(v).add(e);
				return Segment.getFarthestFeature(p1, p2, n, transform);
			}
		}
	}

	/* (non-Javadoc)
	 * @see org.dyn4j.geometry.Shape#project(org.dyn4j.geometry.Vector2, org.dyn4j.geometry.Transform)
	 */
	@Override
	public Interval project(Vector2 n, Transform transform) {
		// get the world space farthest point
		Vector2 p1 = this.getFarthestPoint(n, transform);
		// get the center in world space
		Vector2 center = transform.getTransformed(this.center);
		// project the center onto the axis
		double c = center.dot(n);
		// project the point onto the axis
		double d = p1.dot(n);
		// get the interval along the axis
		return new Interval(2 * c - d, d);
	}

	/* (non-Javadoc)
	 * @see org.dyn4j.geometry.Shape#createAABB(org.dyn4j.geometry.Transform)
	 */
	@Override
	public AABB createAABB(Transform transform) {
		Interval x = this.project(Vector2.X_AXIS, transform);
		Interval y = this.project(Vector2.Y_AXIS, transform);
		
		return new AABB(x.getMin(), y.getMin(), x.getMax(), y.getMax());
	}

	/* (non-Javadoc)
	 * @see org.dyn4j.geometry.Shape#createMass(double)
	 */
	@Override
	public Mass createMass(double density) {
		// the mass of a capsule is the mass of the rectangular section plus the mass
		// of two half circles (really just one circle)
		
		double h = this.capRadius * 2.0;
		double w = this.length - h;
		double r2 = this.capRadius * this.capRadius;
		
		// compute the rectangular area
		double ra = w * h;
		// compuate the circle area
		double ca = r2 * Math.PI;
		double rm = density * ra;
		double cm = density * ca;
		double m = rm + cm;
		
		// the inertia is slightly different. Its the inertia of the rectangular
		// region plus the inertia of half a circle moved from the center
		double d = w * 0.5;
		// parallel axis theorem I2 = Ic + m * d^2
		double cI = 0.5 * cm * r2 + cm * d * d;
		double rI = rm * (h * h + w * w) / 12.0;
		// add the rectangular inertia and cicle inertia
		double I = rI + cI;
		
		return new Mass(this.center, m, I);
	}

	/* (non-Javadoc)
	 * @see org.dyn4j.geometry.Shape#getRadius(org.dyn4j.geometry.Vector2)
	 */
	@Override
	public double getRadius(Vector2 center) {
		return this.radius + this.center.distance(center);
	}

	/* (non-Javadoc)
	 * @see org.dyn4j.geometry.Shape#contains(org.dyn4j.geometry.Vector2, org.dyn4j.geometry.Transform)
	 */
	@Override
	public boolean contains(Vector2 point, Transform transform) {
		// a capsule is just a radially expanded line segment
		Vector2 p = Segment.getPointOnSegmentClosestToPoint(point, transform.getTransformed(this.foci[0]), transform.getTransformed(this.foci[1]));
		double r2 = this.capRadius * this.capRadius;
		double d2 = p.distanceSquared(point);
		return d2 <= r2;
	}

	/* (non-Javadoc)
	 * @see org.dyn4j.geometry.AbstractShape#rotate(double, double, double)
	 */
	@Override
	public void rotate(double theta, double x, double y) {
		super.rotate(theta, x, y);
		// rotate the foci
		this.foci[0].rotate(theta, x, y);
		this.foci[1].rotate(theta, x, y);
		// rotate the local x-axis
		this.localXAxis.rotate(theta);
	}
	
	/* (non-Javadoc)
	 * @see org.dyn4j.geometry.AbstractShape#translate(double, double)
	 */
	@Override
	public void translate(double x, double y) {
		super.translate(x, y);
		// translate the foci
		this.foci[0].add(x, y);
		this.foci[1].add(x, y);
	}

	/**
	 * Returns the rotation about the local center in radians.
	 * @return double the rotation in radians
	 */
	public double getRotation() {
		return Vector2.X_AXIS.getAngleBetween(this.localXAxis);
	}
	
	/**
	 * Returns the length of the capsule.
	 * <p>
	 * The length is the largest dimension of the capsule's
	 * bounding rectangle.
	 * @return double
	 */
	public double getLength() {
		return this.length;
	}
	
	/**
	 * Returns the end cap radius.
	 * @return double
	 */
	public double getCapRadius() {
		return this.capRadius;
	}
}

